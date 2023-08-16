package view

type Style struct {
	Fg       Color
	Bg       Color
	Modifier Modifier
}
type (
	Attr     uint16
	Modifier uint
	Color    int
)

const (
	ColorClear   Color = -1
	ColorBlack   Color = 0
	ColorRed     Color = 1
	ColorGreen   Color = 2
	ColorYellow  Color = 3
	ColorBlue    Color = 4
	ColorMagenta Color = 5
	ColorCyan    Color = 6
	ColorWhite   Color = 7
)

const (
	ModifierClear Modifier = 0
	ModifierBold  Modifier = 1 << (iota + 9)
	ModifierUnderline
	ModifierItalic
	ModifierReverse
)

var StyleClear = Style{Fg: ColorClear, Bg: ColorClear, Modifier: ModifierClear}
