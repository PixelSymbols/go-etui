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

func newPage() *Page {
	page := components.NewPage(core.NewItemContainer())
	return &Page{
		Text: *components.NewText(page),
	}
}
func NewPage() IPage {
	return newPage()
}

type (
	TUI struct {
		*Page
		components.Pager
	}
	ITUI interface {
		IPage
		components.IPager
	}
)

func NewTUI() ITUI {
	container := core.NewItemContainer()
	page := newPage()
	pager := components.NewPager(container)
	pager.AddPage(page)
	return &TUI{
		Page:  page,
		Pager: *pager,
	}
}

func main() {
	tui := NewTUI()
	tui.AddText("hello world")
	page := NewPage()
	page.AddText("second page")
	tui.AddPage(page)
}
