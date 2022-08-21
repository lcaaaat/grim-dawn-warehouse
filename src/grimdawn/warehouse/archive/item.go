package archive

import "bufio"

const MaxItemNumber = 1024

type Item struct {
	baseName      string
	prefixName    string
	suffixName    string
	modifierName  string
	transmuteName string
	seed          uint32
	relicName     string
	relicBonus    string
	relicSeed     uint32
	augmentName   string
	unknown       uint32
	augmentSeed   uint32
	var1          uint32
	stackCount    uint32
	xOffset       float32
	yOffset       float32
}

func (item *Item) Load(arc *Archive, page *Page, reader *bufio.Reader) {
	item.baseName = arc.ReadStr(reader)
	item.prefixName = arc.ReadStr(reader)
	item.suffixName = arc.ReadStr(reader)
	item.modifierName = arc.ReadStr(reader)
	item.transmuteName = arc.ReadStr(reader)
	item.seed = arc.ReadUint32MaskAndUpdateKey(reader)
	item.relicName = arc.ReadStr(reader)
	item.relicBonus = arc.ReadStr(reader)
	item.relicSeed = arc.ReadUint32MaskAndUpdateKey(reader)
	item.augmentName = arc.ReadStr(reader)
	item.unknown = arc.ReadUint32MaskAndUpdateKey(reader)
	item.augmentSeed = arc.ReadUint32MaskAndUpdateKey(reader)
	item.var1 = arc.ReadUint32MaskAndUpdateKey(reader)
	item.stackCount = arc.ReadUint32MaskAndUpdateKey(reader)
	item.xOffset = arc.ReadFloat32(reader)
	item.yOffset = arc.ReadFloat32(reader)
}
