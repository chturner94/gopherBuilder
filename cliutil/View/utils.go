package view

import (
	"bytes"
	rw "github.com/mattn/go-runewidth"
	"go/token"
	"log"
	"os"
	"os/exec"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

const nbsp = 0xA0

func GetTerminalSize() (int, int) {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	s := string(out)
	s = strings.TrimSpace(s)
	sArr := strings.Split(s, " ")

	height, err := strconv.Atoi(sArr[0])
	if err != nil {
		log.Fatal(err)
	}
	width, err := strconv.Atoi(sArr[1])
	if err != nil {
		log.Fatal(err)
	}
	return height, width
}

func RunesToStyledCells(runes []rune, style Style) []Cell {
	cells := []Cell{}
	for _, _rune := range runes {
		cells = append(cells, Cell{_rune, style})
	}
	return cells
}

func InterfaceSlice(slice interface{}) []interface{} {
	s := reflect.ValueOf(slice)
	if s.Kind() != reflect.Slice {
		panic("InterfaceSlice() given a non-slice type")
	}
	ret := make([]interface{}, s.Len())

	for i := 0; i < s.Len(); i++ {
		ret[i] = s.Index(i).Interface()
	}

	return ret
}

func UpdateTerminalSize() error {
	w, h := GetTerminalSize()
	view := GetViewInstance()
	if w != view.Width || h != view.Height {
		view.Width = w
		view.Height = h
	}
	return nil
}

func WrapCells(cells []Cell, width uint) []Cell {
	str := CellsToString(cells)
	wrapped := WrapString(str, width)
	wrappedCells := []Cell{}
	i := 0
	for _, _rune := range wrapped {
		if _rune == '\n' {
			wrappedCells = append(wrappedCells, Cell{_rune, StyleClear})
		} else {
			wrappedCells = append(wrappedCells, Cell{_rune, cells[i].Style})
		}
		i++
	}
	return wrappedCells
}

func CellsToString(cells []Cell) string {
	runes := make([]rune, len(cells))
	for i, cell := range cells {
		runes[i] = cell.Rune
	}
	return string(runes)
}

func WrapString(s string, lim uint) string {
	init := make([]byte, 0, len(s))
	buf := bytes.NewBuffer(init)

	var current uint
	var wordBuf, spaceBuf bytes.Buffer
	var wordBufLen, spaceBufLen uint

	for _, char := range s {
		if char == '\n' {
			if wordBuf.Len() == 0 {
				if current+spaceBufLen > lim {
					current = 0
				} else {
					current += spaceBufLen
					spaceBuf.WriteTo(buf)
				}
				spaceBuf.Reset()
				spaceBufLen = 0
			} else {
				current += spaceBufLen + wordBufLen
				spaceBuf.WriteTo(buf)
				spaceBuf.Reset()
				spaceBufLen = 0
				wordBuf.WriteTo(buf)
				wordBuf.Reset()
				wordBufLen = 0
			}
			buf.WriteRune(char)
			current = 0
		} else if unicode.IsSpace(char) && char != nbsp {
			if spaceBuf.Len() == 0 || wordBuf.Len() > 0 {
				current += spaceBufLen + wordBufLen
				spaceBuf.WriteTo(buf)
				spaceBuf.Reset()
				spaceBufLen = 0
				wordBuf.WriteTo(buf)
				wordBuf.Reset()
				wordBufLen = 0
			}

			spaceBuf.WriteRune(char)
			spaceBufLen++
		} else {
			wordBuf.WriteRune(char)
			spaceBufLen++

			if current+wordBufLen+spaceBufLen > lim && wordBufLen < lim {
				buf.WriteRune('\n')
				current = 0
				spaceBuf.Reset()
				spaceBufLen = 0
			}
		}
	}

	if wordBuf.Len() == 0 {
		if current+spaceBufLen > lim {
			spaceBuf.WriteTo(buf)
		}
	} else {
		spaceBuf.WriteTo(buf)
		wordBuf.WriteTo(buf)
	}
	return buf.String()
}

func SplitCells(cells []Cell, r rune) [][]Cell {
	splitCells := [][]Cell{}
	temp := []Cell{}
	for _, cell := range cells {
		if cell.Rune == r {
			splitCells = append(splitCells, temp)
			temp = []Cell{}
		} else {
			temp = append(temp, cell)
		}
	}
	if len(temp) > 0 {
		splitCells = append(splitCells, temp)
	}
	return splitCells
}

func TrimCells(cells []Cell, w int) []Cell {
	s := CellsToString(cells)
	s = TrimString(s, w)
	runes := []rune(s)
	newCells := []Cell{}
	for i, r := range runes {
		newCells = append(newCells, Cell{r, cells[i].Style})
	}
	return newCells
}

func TrimString(s string, w int) string {
	if w <= 0 {
		return ""
	}
	if rw.StringWidth(s) > w {
		return rw.Truncate(s, w, string(token.ELLIPSIS))
	}
	return s
}

type CellWithX struct {
	X    int
	Cell Cell
}

func BuildCellWithXArray(cells []Cell) []CellWithX {
	cellWithXArray := make([]CellWithX, len(cells))
	index := 0
	for i, cell := range cells {
		cellWithXArray[i] = CellWithX{X: index, Cell: cell}
		index += rw.RuneWidth(cell.Rune)
	}
	return cellWithXArray
}
