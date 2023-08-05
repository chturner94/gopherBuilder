package style

import "github.com/chturner94/gopherBuilder/cliutil/View"

var BaseColors = []Color{
	ColorBlack,
	ColorRed,
	ColorGreen,
	ColorYellow,
	ColorBlue,
	ColorMagenta,
	ColorCyan,
	ColorWhite,
}

var BaseStyles = []Style{
	view.NewStyle(ColorBlack),
	view.NewStyle(ColorRed),
	view.NewStyle(ColorGreen),
	view.NewStyle(ColorYellow),
	view.NewStyle(ColorBlue),
	view.NewStyle(ColorMagenta),
	view.NewStyle(ColorCyan),
	view.NewStyle(ColorWhite),
}

// BaseTheme is where all the mappings for modules and the corresponding Type, which defines the expected properties
// for that module are defined. For instance, the Text module is mapped to the TextTheme struct, which expects a single
// property, Text, which holds the property for the color of the text.
//
// To add a new module, first create an exported struct
// with the required properties (types set to Color or Style), and then add it to the BaseTheme struct. Naming convention
// should be to use the name of the module followed by theme for the struct name, and using only the module name for the
// property name in BaseTheme.
type BaseTheme struct {
	Default Style
	// Put module themes here
	Module ModuleTheme
	Text   TextTheme
}

type ModuleTheme struct {
	Title  Style
	Border Style
}

type TextTheme struct {
	Text Style
}

var Theme = BaseTheme{
	Default: view.NewStyle(ColorWhite),

	Module: ModuleTheme{
		Title:  view.NewStyle(ColorWhite),
		Border: view.NewStyle(ColorWhite),
	},
	Text: TextTheme{
		Text: view.NewStyle(ColorWhite),
	},
}
