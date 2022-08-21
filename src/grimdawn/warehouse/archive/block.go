package archive

import (
	"bufio"
	"errors"
)

type Block struct {
	len uint32
}

func (block *Block) ReadBlockStart(arc *Archive, reader *bufio.Reader) uint32 {
	ret := arc.ReadUint32MaskAndUpdateKey(reader)
	block.len = arc.ReadUint32Mask(reader)
	return ret
}

func (block *Block) ReadBlockEnd(arc *Archive, reader *bufio.Reader) error {
	if arc.ReadUint32Mask(reader) != 0 {
		return errors.New("unexpected gst file format")
	}
	return nil
}
