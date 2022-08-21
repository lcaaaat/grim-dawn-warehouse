package archive

import "bufio"

const MaxPageNumber = 16

type Page struct {
	block    Block
	width    uint32
	height   uint32
	itemSize uint32
	items    [MaxItemNumber]Item
}

func (page *Page) Load(arc *Archive, reader *bufio.Reader) error {
	page.block.ReadBlockStart(arc, reader)
	page.width = arc.ReadUint32MaskAndUpdateKey(reader)
	page.height = arc.ReadUint32MaskAndUpdateKey(reader)
	page.itemSize = arc.ReadUint32MaskAndUpdateKey(reader)
	for i := uint32(0); i < page.itemSize; i++ {
		page.items[i].Load(arc, page, reader)
	}
	return page.block.ReadBlockEnd(arc, reader)
}
