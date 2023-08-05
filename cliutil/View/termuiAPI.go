package view

import "github.com/chturner94/gopherBuilder/cliutil/View/style"

func NewStyle(fg style.Color, args ...interface{}) style.Style {
	bg := style.ColorClear
	modifier := style.ModifierClear
	if len(args) >= 1 {
		bg = args[0].(style.Color)
	}
	if len(args) == 2 {
		modifier = args[1].(style.Modifier)
	}
	return style.Style{
		fg,
		bg,
		modifier,
	}
}
