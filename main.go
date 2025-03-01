package main

import (
	"github.com/Pixelsymbols/go-etui/components"
	"github.com/Pixelsymbols/go-etui/core"
)

type (
	Page struct {
		components.Page
		components.Text
	}
	IPage interface {
		components.IPage
		components.IText
	}
)

func NewPage() *Page {
	page := components.NewPage(core.NewItemContainer())
	return &Page{
		Text: *components.NewText(page),
	}
}

type (
	TUI struct {
		Page
		components.Pager
	}
	ITUI interface {
		IPage
		components.IPager
	}
)

func NewTUI() ITUI {
	return &TUI{
		Page: *NewPage(),
	}
}

func main() {
	tui := NewTUI()
	tui.AddText("hello world")
}
