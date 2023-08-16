package view

import (
	"image"
	"sync"
)

type canvasItemType uint

const (
	col canvasItemType = iota
	row
)

var viewInstance *View
var once sync.Once

type View struct {
	canvas  []Canvas
	Width   int
	Height  int
	offsetX int
	offsetY int
	cursorX int
	cursorY int
	zoom    float64
}

func GetViewInstance() *View {
	once.Do(func() {
		width, height := GetTerminalSize()
		viewInstance = &View{
			Width:  width,
			Height: height,
		}
	})
	return viewInstance
}

type Canvas struct {
	Module
	Objects []*CanvasObject
}

type CanvasObject struct {
	rowColumnType canvasItemType
	XRatio        float64
	YRatio        float64
	WidthRatio    float64
	HeightRatio   float64
	ratio         float64
	Entry         interface{}
	IsLeaf        bool
	Cell          [][]Cell
}

func NewCanvas() *Canvas {
	c := &Canvas{
		Module: *NewModule(),
	}
	c.Border = false
	return c
}

func NewCol(ratio float64, i ...interface{}) CanvasObject {
	_, ok := i[0].(Drawable)
	entry := i[0]
	if !ok {
		entry = i
	}
	return CanvasObject{
		rowColumnType: col,
		Entry:         entry,
		IsLeaf:        ok,
		ratio:         ratio,
	}
}

func NewRow(ratio float64, i ...interface{}) CanvasObject {
	_, ok := i[0].(Drawable)
	entry := i[0]
	if !ok {
		entry = i
	}
	return CanvasObject{
		rowColumnType: row,
		Entry:         entry,
		IsLeaf:        ok,
		ratio:         ratio,
	}
}

func (self *Canvas) Set(module ...interface{}) {
	entry := CanvasObject{
		rowColumnType: row,
		Entry:         module,
		IsLeaf:        false,
		ratio:         1.0,
	}
	self.setHelper(entry, 1.0, 1.0)
}

func (self *Canvas) setHelper(object CanvasObject, parentWidthRatio, parentHeightRatio float64) {
	var HeightRatio float64
	var WidthRatio float64
	switch object.rowColumnType {
	case col:
		HeightRatio = 1.0
		WidthRatio = object.ratio
	case row:
		HeightRatio = object.ratio
		WidthRatio = 1.0
	}
	object.WidthRatio = parentWidthRatio * WidthRatio
	object.HeightRatio = parentHeightRatio * HeightRatio

	if object.IsLeaf {
		self.Objects = append(self.Objects, &object)
	} else {
		XRatio := 0.0
		YRatio := 0.0
		cols := false
		rows := false

		children := InterfaceSlice(object.Entry)

		for i := 0; i < len(children); i++ {
			if children[i] == nil {
				continue
			}
			child, _ := children[i].(CanvasObject)

			child.XRatio = object.XRatio + (object.WidthRatio * XRatio)
			child.YRatio = object.YRatio + (object.HeightRatio * YRatio)

			switch child.rowColumnType {
			case col:
				cols = true
				XRatio += child.ratio
				if rows {
					object.HeightRatio /= 2
				}
			case row:
				rows = true
				YRatio += child.ratio
				if cols {
					object.WidthRatio /= 2
				}
			}
			self.setHelper(child, object.WidthRatio, object.HeightRatio)
		}
	}
}

func (self *Canvas) Draw(buf *Buffer) {
	width := float64(self.Dx()) + 1
	height := float64(self.Dy()) + 1

	for _, object := range self.Objects {
		entry, _ := object.Entry.(Drawable)

		x := int(width*object.XRatio) + self.Min.X
		y := int(height*object.YRatio) + self.Min.Y
		w := int(width * object.WidthRatio)
		h := int(height * object.HeightRatio)

		if x+w > self.Dx() {
			w--
		}
		if y+h > self.Dy() {
			h--
		}

		entry.SetRect(x, y, x+w, y+h)

		entry.Lock()
		entry.Draw(buf)
		entry.Unlock()
	}
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

func (self *Buffer) SetString(s string, style Style, p image.Point) {
	runes := []rune(s)
	x := 0
	for _, char := range runes {
		self.SetCell(Cell{char, style}, image.Pt(p.X+x, p.Y))
		x += charWidth(char)
	}
}

// Replace with a golang.org/x/text/Width implementation
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
	Style Style
}

var CellEmpty = Cell{
	Rune:  ' ',
	Style: StyleClear,
}

func NewCell(rune rune, args ...interface{}) Cell {
	updateStyle := StyleClear
	if len(args) == 1 {
		updateStyle = args[0].(Style)
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
