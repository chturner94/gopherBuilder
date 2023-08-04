package View

type (
	Attr uint16
)

type Canvas [][]Cell

func NewCanvas(width, height int) Canvas {
	canvas := make(Canvas, width)
	for i := range canvas {
		canvas[i] = make([]Cell, height)
	}
	return canvas
}

type module struct {
}

type text struct {
	posX   int
	posY   int
	text   []rune
	canvas []Cell
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
