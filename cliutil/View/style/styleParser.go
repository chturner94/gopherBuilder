package style

import (
	view "github.com/chturner94/gopherBuilder/cliutil/View"
	utils "github.com/chturner94/gopherBuilder/cliutil/View/utils"
	"strings"
)

const (
	tokenFg       = "fg"
	tokenBg       = "bg"
	tokenModifier = "mod"

	tokenItemSeparator   = ","
	tokenValueSeparator  = ":"
	tokenBeginStyledText = '['
	tokenEndStyledText   = ']'
	tokenBeginStyle      = '{'
	tokenEndStyle        = '}'
)

type parserState uint

const (
	parserStateDefault parserState = iota
	parserStateStyleItems
	parserStateStyledText
)

var StyledParserColorMap = map[string]Color{
	"black":   ColorBlack,
	"red":     ColorRed,
	"green":   ColorGreen,
	"yellow":  ColorYellow,
	"blue":    ColorBlue,
	"magenta": ColorMagenta,
	"cyan":    ColorCyan,
	"white":   ColorWhite,
}

var modifierMap = map[string]Modifier{
	"bold":      ModifierBold,
	"underline": ModifierUnderline,
	"italic":    ModifierItalic,
	"reverse":   ModifierReverse,
}

func readStyle(runes []rune, defaultStyle Style) Style {
	style := defaultStyle
	split := strings.Split(string(runes), tokenItemSeparator)
	for _, item := range split {
		pair := strings.Split(item, tokenValueSeparator)
		if len(pair) == 2 {
			switch pair[0] {
			case tokenFg:
				style.Fg = StyledParserColorMap[pair[1]]
			case tokenBg:
				style.Bg = StyledParserColorMap[pair[1]]
			case tokenModifier:
				style.Modifier = modifierMap[pair[1]]
			}
		}
	}
	return style
}

func ParseStyles(s string, defaultStyle Style) []view.Cell {
	var cells []view.Cell
	runes := []rune(s)
	state := parserStateDefault
	var styledText []rune
	var styledItems []rune
	squareCount := 0

	reset := func() {
		state = parserStateDefault
		styledText = []rune{}
		styledItems = []rune{}
		squareCount = 0
	}

	rollback := func() {
		cells = append(cells, utils.RunesToStyledCells(styledText, defaultStyle)...)
		cells = append(cells, utils.RunesToStyledCells(styledItems, defaultStyle)...)
		reset()
	}

	chop := func(s []rune) []rune {
		return s[1 : len(s)-1]
	}

	for i, _rune := range runes {
		switch state {
		case parserStateDefault:
			if _rune == tokenBeginStyledText {
				state = parserStateStyledText
				squareCount = 1
				styledText = append(styledText, _rune)
			} else {
				cells = append(cells, view.Cell{Rune: _rune, Style: defaultStyle})
			}
		case parserStateStyledText:
			switch {
			case squareCount == 0:
				switch _rune {
				case tokenBeginStyle:
					state = parserStateStyleItems
					styledItems = append(styledItems, _rune)
				default:
					rollback()
					switch _rune {
					case tokenBeginStyledText:
						state = parserStateStyledText
						squareCount = 1
						styledItems = append(styledItems, _rune)
					default:
						cells = append(cells, view.Cell{Rune: _rune, Style: defaultStyle})
					}
				}
			case len(runes) == i+1:
				rollback()
				styledText = append(styledText, _rune)
			case _rune == tokenBeginStyledText:
				squareCount++
				styledText = append(styledText, _rune)
			case _rune == tokenEndStyledText:
				squareCount--
				styledText = append(styledText, _rune)
			default:
				styledText = append(styledText, _rune)
			}
		case parserStateStyleItems:
			styledItems = append(styledItems, _rune)
			if _rune == tokenEndStyle {
				style := readStyle(chop(styledItems), defaultStyle)
				cells = append(cells, utils.RunesToStyledCells(chop(styledText), style)...)
				reset()
			} else if len(runes) == i+1 {
				rollback()
			}
		}
	}

	return cells
}
