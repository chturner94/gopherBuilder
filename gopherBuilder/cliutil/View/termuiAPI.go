package View

type View struct {
	Canvas  Canvas
	Modules []Module
	Width   int
	Height  int
	OffsetX int
	OffsetY int
	Zoom    float64
	Delta   float64
}

func NewView() *View {
	v := View{
		Modules: make([]Module, 0),
	}
	s.Canvas = NewCanvas(10, 10)
	return &v
}

type Cell struct {
	Fg Attr // Foreground color
	Bg Attr // Background color
	Ch rune // Unicode character to draw
}

type Event struct {
	Type EventType
	// Key Key
	Ch rune
	// KeyMod KeyMod
	Err    error
	MouseX int
	MouseY int
}

type (
	EventType uint8
)
