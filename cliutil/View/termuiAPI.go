package View

import (
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

type View struct {
	Canvas  Canvas
	Width   int
	Height  int
	OffsetX int
	OffsetY int
	CursorX int
	CursorY int
	Zoom    float64
}
type Canvas struct {
	Height  int
	Width   int
	Modules []Module
	Cell    [][]Cell
}

type Module struct {
	PosX int
	PosY int
}

type Cell struct {
	Fg Attr // Foreground color
	Bg Attr // Background color
	Ch rune // An unicode character to draw
}

func (v *View) Init() {
	v.Height, v.Width = getTerminalSize()
	v.Canvas.Height = v.Height
	v.Canvas.Width = v.Width
	v.Canvas.Cell = make([][]Cell, v.Height)
	for i := range v.Canvas.Cell {
		v.Canvas.Cell[i] = make([]Cell, v.Width)
	}
}

func (v *View) Render() {
	// Clear the screen
	print("\033[H\033[2J")
	// Render the canvas
	for y := 0; y < v.Height; y++ {
		for x := 0; x < v.Width; x++ {
			cell := v.Canvas.Cell[y][x]

			print(cell.Ch)
		}
		print("\n")
	}

	// Keep Rendering until application is closed
	for {
		print("\033[H\033[2J")
		for y := 0; y < v.Height; y++ {
			for x := 0; x < v.Width; x++ {
				cell := v.Canvas.Cell[y][x]
				print(cell.Ch)
			}
			print("\n")
		}
	}
	if shouldExit() {
		os.Exit(0)
	}

}

func shouldExit() bool {
	return false
}

func getTerminalSize() (int, int) {
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
