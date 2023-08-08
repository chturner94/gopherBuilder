package modules

import (
	view "github.com/chturner94/gopherBuilder/cliutil/View"
	"github.com/chturner94/gopherBuilder/cliutil/View/style"
	"github.com/chturner94/gopherBuilder/cliutil/View/utils"
)

type Text struct {
	Module
	Text      string
	TextStyle style.Style
	WrapText  bool
}

func NewText() *Text {
	return &Text{
		Module:    *NewModule(),
		TextStyle: style.Theme.Text.Text,
		WrapText:  true,
	}
}

func (self *Text) Draw(buf *view.Buffer) {
	self.Module.Draw(buf)

	cells := style.ParseStyles(self.Text, self.TextStyle)
	if self.WrapText {
		cells = utils.WrapCells(cells, uint(self.Inner.Dx()))
	}

	rows := utils.SplitCells(cells, '\n')

	for y, row := range rows {
		if y+self.Inner.Min.Y >= self.Inner.Max.Y {
			break
		}
		row = TrimCell
	}
}
