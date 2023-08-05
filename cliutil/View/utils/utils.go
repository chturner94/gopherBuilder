package utils

import (
	view "github.com/chturner94/gopherBuilder/cliutil/View"
	"github.com/chturner94/gopherBuilder/cliutil/View/style"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

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

func RunesToStyledCells(runes []rune, style style.Style) []view.Cell {
	cells := []view.Cell{}
	for _, _rune := range runes {
		cells = append(cells, view.Cell{_rune, style})
	}
	return cells
}
