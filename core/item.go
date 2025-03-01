package core

type Item struct {
}
type IItem interface {
}

type ItemContainer struct {
	items []IItem
}
type IItemContainer interface {
	GetItems() []IItem
}

func (c *ItemContainer) AddItem(item IItem) {
	c.items = append(c.items, item)
}

func (i *ItemContainer) GetItems() []IItem {
	return i.items
}

func NewItemContainer(size ...int) *ItemContainer {
	return &ItemContainer{
		items: make([]IItem, max(4, size[max(0, len(size)-1)])),
	}
}
