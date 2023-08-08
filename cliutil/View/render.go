package view

import (
	"github.com/gdamore/tcell/termbox"
	"image"
	"sync"
)

type Drawable interface {
	GetRect() image.Rectangle
	SetRect(int, int, int, int)
	Draw(*Buffer)
	sync.Locker
}

func Render(objects ...Drawable) {
	for _, obj := range objects {
		buf := NewBuffer(obj.GetRect())
		obj.Lock()
		obj.Draw(buf)
		obj.Unlock()
		for point, cell := range buf.CellMap {
			if point.In(buf.Rectangle) {
				termbox.SetCell(
					point.X, point.Y,
					cell.Rune,
					termbox.Attribute(cell.Style.Fg+1)|termbox.Attribute(cell.Style.Modifier), termbox.Attribute(cell.Style.Bg+1),
				)
			}
		}
	}
	termbox.Flush()
}
