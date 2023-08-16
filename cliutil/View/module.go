package view

import (
	"image"
	"sync"
)

type Module struct {
	Border      bool
	BorderStyle Style

	BorderLeft, BorderRight, BorderTop, BorderBottom     bool
	PaddingLeft, PaddingRight, PaddingTop, PaddingBottom int

	image.Rectangle
	Inner image.Rectangle

	Title      string
	TitleStyle Style

	data interface{}
	sync.Mutex
}

func NewModule() *Module {
	return &Module{
		Border:       true,
		BorderStyle:  Theme.Module.Border,
		BorderLeft:   true,
		BorderRight:  true,
		BorderTop:    true,
		BorderBottom: true,

		TitleStyle: Theme.Module.Title,
	}
}

func (m *Module) drawBorder(buf *Buffer) {
	verticalCell := Cell{VERTICAL_LINE, m.BorderStyle}
	horizontalCell := Cell{HORIZONTAL_LINE, m.BorderStyle}

	if m.BorderTop {
		buf.Fill(horizontalCell, image.Rect(m.Min.X, m.Min.Y, m.Max.X, m.Min.Y+1))
	}
	if m.BorderBottom {
		buf.Fill(horizontalCell, image.Rect(m.Min.X, m.Max.Y-1, m.Max.X, m.Max.Y))
	}
	if m.BorderLeft {
		buf.Fill(verticalCell, image.Rect(m.Min.X, m.Min.Y, m.Min.X+1, m.Max.Y))
	}
	if m.BorderRight {
		buf.Fill(verticalCell, image.Rect(m.Max.X-1, m.Min.Y, m.Max.X, m.Max.Y))
	}

	//corners
	if m.BorderTop && m.BorderLeft {
		buf.SetCell(Cell{TOP_LEFT_CORNER, m.BorderStyle}, image.Pt(m.Min.X, m.Min.Y))
	}
	if m.BorderTop && m.BorderRight {
		buf.SetCell(Cell{TOP_RIGHT_CORNER, m.BorderStyle}, image.Pt(m.Max.X-1, m.Min.Y))
	}
	if m.BorderBottom && m.BorderLeft {
		buf.SetCell(Cell{BOTTOM_LEFT_CORNER, m.BorderStyle}, image.Pt(m.Min.X, m.Max.Y-1))
	}
	if m.BorderBottom && m.BorderRight {
		buf.SetCell(Cell{BOTTOM_RIGHT_CORNER, m.BorderStyle}, image.Pt(m.Max.X-1, m.Max.Y-1))
	}
}

func (m *Module) SetRect(x1, y1, x2, y2 int) {
	m.Lock()
	defer m.Unlock()
	m.Rectangle = image.Rect(x1, y1, x2, y2)
	m.Inner = image.Rect(
		m.Min.X+m.PaddingLeft+1,
		m.Min.Y+m.PaddingTop+1,
		m.Max.X-m.PaddingRight-1,
		m.Max.Y-m.PaddingBottom-1,
	)
}

func (self *Module) GetRect() image.Rectangle {
	return self.Rectangle
}

func (m *Module) Draw(buf *Buffer) {
	m.Lock()
	defer m.Unlock()
	if m.Border {
		m.drawBorder(buf)
	}

	if m.Title != "" {
		buf.SetString(m.Title, m.TitleStyle, image.Pt(m.Min.X+1, m.Min.Y))
	}
}

//func (m *Module) Draw(canvas *view.Buffer) {
//	m.Lock()
//	defer m.Unlock()
//	startX := m.Inner.Min.X
//	startY := m.Inner.Min.Y
//
//	for y := startY; y < startY+m.Inner.Dy(); y++ {
//		for x := startX; x < startX+m.Inner.Dx(); x++ {
//			cell := &canvas.Cell[y][x]
//
//			text := m.data.(string)
//			cell.Rune = rune(text[x-startX])
//		}
//	}
//}

func (m *Module) SetData(data interface{}) {
	m.data = data
}
