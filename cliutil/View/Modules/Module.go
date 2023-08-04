package Modules

import (
	View "github.com/chturner94/gopherBuilder/cliutil/View"
	"image"
	"sync"
)

type Module struct {
	Border                                           bool
	BorderStyle                                      View.Style
	BorderLeft, BorderRight, BorderTop, BorderBottom bool

	PaddingLeft, PaddingRight, PaddingTop, PaddingBottom int

	image.Rectangle
	Inner image.Rectangle

	Title      string
	TitleStyle View.Style

	sync.Mutex
}

func NewModule() *Module {
	return &Module{
		Border:       true,
		BorderLeft:   true,
		BorderRight:  true,
		BorderTop:    true,
		BorderBottom: true,
	}
}

const (
	TOP_LEFT     = '┌'
	TOP_RIGHT    = '┐'
	BOTTOM_LEFT  = '└'
	BOTTOM_RIGHT = '┘'

	HORIZONTAL = '─'
	VERTICAL   = '│'

	TOP_T    = '┬'
	BOTTOM_T = '┴'
	LEFT_T   = '├'
	RIGHT_T  = '┤'

	QUOTA_LEFT  = '«'
	QUOTA_RIGHT = '»'

	VERTICAL_DASH   = '┊'
	HORIZONTAL_DASH = '┄'
)
