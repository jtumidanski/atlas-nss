package item

type Model struct {
	id       uint32
	shopId   uint32
	itemId   uint32
	price    uint32
	pitch    uint32
	position uint32
}

func (m Model) ItemId() uint32 {
	return m.itemId
}

func (m Model) Price() uint32 {
	return m.price
}

func (m Model) Pitch() uint32 {
	return m.pitch
}

func (m Model) Position() uint32 {
	return m.position
}
