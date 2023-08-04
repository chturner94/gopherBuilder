package View

type Style struct {
	Fg       Color
	Bg       Color
	Modifier Modifier
}

type Color int

const (
	ColorBlack Color = iota
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

type Modifier uint

const (
	AttrBold Modifier = 1 << (iota + 9)
	AttrUnderline
	AttrItalic
	AttributeReverse
	AttrClear Modifier = 0
)
