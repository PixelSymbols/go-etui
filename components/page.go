package components

import "github.com/Pixelsymbols/go-etui/core"

type (
	Pager struct {
		container *core.ItemContainer
	}
	IPager interface {
		GetPages() *core.ItemContainer
		AddPage(page IPage)
	}
	Page struct {
		core.Item
		core.ItemContainer
	}
	IPage interface {
		GetItems() []core.IItem
		core.IItem
	}
)

func (p *Pager) AddPage(page IPage) {
	p.container.AddItem(page)
}
func (p *Pager) GetPages() *core.ItemContainer {
	return p.container
}
func NewPager(container *core.ItemContainer) *Pager {
	return &Pager{
		container: container,
	}
}

func NewPage(container *core.ItemContainer) IPage {
	return &Page{
		ItemContainer: *core.NewItemContainer(),
	}
}
