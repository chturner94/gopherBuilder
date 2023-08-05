package view

import (
	"image"

	"github.com/chturner94/gopherBuilder/cliutil/View/modules"
	"github.com/chturner94/gopherBuilder/cliutil/View/style"
)

type viewItemType uint

const (
	col viewItemType = iota
	row
)

type View struct {
	rowColumnType viewItemType
	canvas        []Canvas
	width         int
	height        int
	offsetX       int
	offsetY       int
	cursorX       int
	cursorY       int
	zoom          float64
}

type Canvas struct {
	height  int
	width   int
	modules []modules.Module
	Cell    [][]Cell
}

type Buffer struct {
	image.Rectangle
	CellMap map[image.Point]Cell
}

func NewBuffer(r image.Rectangle) *Buffer {
	buf := &Buffer{
		Rectangle: r,
		CellMap:   make(map[image.Point]Cell),
	}
	buf.Fill(CellEmpty, r)
	return buf
}

func (self *Buffer) GetCell(p image.Point) Cell {
	return self.CellMap[p]
}

func (self *Buffer) SetCell(c Cell, p image.Point) {
	self.CellMap[p] = c
}

func (self *Buffer) Fill(c Cell, r image.Rectangle) {
	for x := r.Min.X; x < r.Max.X; x++ {
		for y := r.Min.Y; y < r.Max.Y; y++ {
			self.SetCell(c, image.Point{x, y})
		}
	}
}

func (self *Buffer) SetString(s string, style style.Style, p image.Point) {
	runes := []rune(s)
	x := 0
	for _, char := range runes {
		self.SetCell(Cell{char, style}, image.Pt(p.X+x, p.Y))
		x += charWidth(char)
	}
}

// Replace with a golang.org/x/text/width implementation
func charWidth(char rune) int {
	if char == '\t' {
		return 4
	}
	return 1
}

/*func NewCanvas() *Canvas {
	c := &Canvas{}
}*/

type Cell struct {
	Rune  rune
	Style style.Style
}

var CellEmpty = Cell{
	Rune:  ' ',
	Style: style.StyleClear,
}

func NewCell(rune rune, args ...interface{}) Cell {
	updateStyle := style.StyleClear
	if len(args) == 1 {
		updateStyle = args[0].(style.Style)
	}
	return Cell{
		Rune:  rune,
		Style: updateStyle,
	}
}

// TODO:
// - New Canvas
// - Draw Border
// - Canvas SetRect
// - Canvas GetRect
// - Work through remaining module methods
// - implement rendering
// - test
