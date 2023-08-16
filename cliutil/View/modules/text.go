package modules

import (
	view "github.com/chturner94/gopherBuilder/cliutil/View"
	"image"
)

type Text struct {
	view.Module
	Text      string
	TextStyle view.Style
	WrapText  bool
}

func NewText() *Text {
	return &Text{
		Module:    *view.NewModule(),
		TextStyle: view.Theme.Text.Text,
		WrapText:  true,
	}
}

func (self *Text) Draw(buf *view.Buffer) {
	self.Module.Draw(buf)

	cells := view.ParseStyles(self.Text, self.TextStyle)
	if self.WrapText {
		cells = view.WrapCells(cells, uint(self.Inner.Dx()))
	}

	rows := view.SplitCells(cells, '\n')

	for y, row := range rows {
		if y+self.Inner.Min.Y >= self.Inner.Max.Y {
			break
		}
		row = view.TrimCells(row, self.Inner.Dx())
		for _, cx := range view.BuildCellWithXArray(row) {
			x, cell := cx.X, cx.Cell
			buf.SetCell(cell, image.Pt(x, y).Add(self.Inner.Min))
		}
	}
}
