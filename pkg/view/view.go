package view

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

type Canvas struct {
	//Module []Module
}
