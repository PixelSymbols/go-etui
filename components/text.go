package components

import "github.com/Pixelsymbols/go-etui/core"

type Text struct {
	core.Item
	container core.IItemContainer
}

type IText interface {
	core.IItem
	AddText(text ...string)
}

func (t *Text) AddText(text ...string) {
	// t.container.AddItem()
}

func NewText(container core.IItemContainer) *Text {
	return &Text{
		container: container,
	}
}
