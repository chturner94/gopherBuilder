package Modules

import (
	View "github.com/chturner94/gopherBuilder/cliutil/View"
)

type Text struct {
	Text      string
	TextStyle View.Style
	WrapText  bool
}

func NewText() *Text {
	return &Text{}
}
