package modules

import (
	view "github.com/chturner94/gopherBuilder/cliutil/View"
	"github.com/chturner94/gopherBuilder/cliutil/View/style"
	"image"
	"sync"
)

type Module struct {
	Border      bool
	BorderStyle style.Style

	BorderLeft, BorderRight, BorderTop, BorderBottom     bool
	PaddingLeft, PaddingRight, PaddingTop, PaddingBottom bool

	image.Rectangle
	Inner image.Rectangle

	Title      string
	TitleStyle style.Style

	data interface{}
	sync.Mutex
}

func NewModule() *Module {
	return &Module{
		Border:       true,
		BorderStyle:  style.Theme.Module.Border,
		BorderLeft:   true,
		BorderRight:  true,
		BorderTop:    true,
		BorderBottom: true,

		TitleStyle: style.Theme.Module.Title,
	}
}

func (m *Module) SetRect(x1, y1, x2, y2 int) {
	m.Lock()
	defer m.Unlock()
	m.Inner = image.Rect(x1, y1, x2, y2)
}

func (m *Module) Draw(canvas *view.Buffer) {
	m.Lock()
	defer m.Unlock()
	startX := m.Inner.Min.X
	startY := m.Inner.Min.Y

	for y := startY; y < startY+m.Inner.Dy(); y++ {
		for x := startX; x < startX+m.Inner.Dx(); x++ {
			cell := &canvas.Cell[y][x]

			text := m.data.(string)
			cell.Rune = rune(text[x-startX])
		}
	}
}

func (m *Module) SetData(data interface{}) {
	m.data = data
}
