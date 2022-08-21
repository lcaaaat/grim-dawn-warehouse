package archive

import (
	"bufio"
	"encoding/binary"
	"errors"
	"github.com/sirupsen/logrus"
	"io"
	"math"
	"strings"
)

type Archive struct {
	key      uint32
	table    [256]uint32
	block    Block
	pageSize uint32
	pages    [MaxPageNumber]Page
}

func (arc *Archive) ReadByteAndUpdateKey(reader *bufio.Reader) byte {
	b, err := reader.ReadByte()
	if err != nil {
		logrus.Fatalf("Read byte failed, cause: %s", err.Error())
	}
	ret := b ^ byte(arc.key)
	arc.key ^= arc.table[b]
	return ret
}

func (arc *Archive) ReadUint32(reader *bufio.Reader) uint32 {
	var arr [4]byte
	_, err := io.ReadFull(reader, arr[:])
	if err != nil {
		logrus.Fatalf("Read int failed, cause: %s", err.Error())
	}

	return binary.LittleEndian.Uint32(arr[:])
}

func (arc *Archive) ReadUint32Mask(reader *bufio.Reader) uint32 {
	ret := arc.ReadUint32(reader)
	ret ^= arc.key
	return ret
}

func (arc *Archive) ReadUint32MaskAndUpdateKey(reader *bufio.Reader) uint32 {
	k := arc.key
	ret := arc.ReadUint32(reader)
	for i := 0; i < 32; i += 8 {
		b := byte(ret >> i)
		k ^= arc.table[b]
	}
	ret ^= arc.key
	arc.key = k
	return ret
}

func (arc *Archive) ReadStr(reader *bufio.Reader) string {
	len := arc.ReadUint32MaskAndUpdateKey(reader)
	var builder strings.Builder
	for i := uint32(0); i < len; i++ {
		b := arc.ReadByteAndUpdateKey(reader)
		builder.WriteByte(b)
	}
	return builder.String()
}

func (arc *Archive) ReadFloat32(reader *bufio.Reader) float32 {
	bits := arc.ReadUint32MaskAndUpdateKey(reader)
	return math.Float32frombits(bits)
}

func (arc *Archive) Load(reader *bufio.Reader) error {
	key := arc.ReadUint32(reader)
	key = 0x55555555 ^ key
	arc.key = key
	for i := 0; i < 256; i++ {
		key = (key >> 1) | (key << 31)
		key *= 39916801
		arc.table[i] = key
	}

	if arc.ReadUint32MaskAndUpdateKey(reader) != 2 {
		return errors.New("unexpected gst file format")
	}

	arc.block.ReadBlockStart(arc, reader)

	if arc.ReadUint32MaskAndUpdateKey(reader) != 5 {
		return errors.New("unexpected gst file format")
	}

	if arc.ReadUint32Mask(reader) != 0 {
		return errors.New("unexpected gst file format")
	}

	if arc.ReadStr(reader) != "" {
		return errors.New("unexpected gst file format")
	}

	if arc.ReadByteAndUpdateKey(reader) != 3 {
		return errors.New("unexpected gst file format")
	}

	arc.pageSize = arc.ReadUint32MaskAndUpdateKey(reader)
	for i := uint32(0); i < arc.pageSize; i++ {
		arc.pages[i].Load(arc, reader)
	}

	err := arc.block.ReadBlockEnd(arc, reader)
	if err != nil {
		return err
	}

	logrus.Infof("Load archive successfully.")
	return nil
}
