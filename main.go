package main

import (
	"github.com/chturner94/gopherBuilder/cliutil/View/modules"
	"github.com/chturner94/gopherBuilder/cliutil/View/style"
	"github.com/gdamore/tcell/termbox"
	"log"
)

func main() {
	if err := Init(); err != nil {
		log.Fatalf("failed to initialize GUI: %v", err)
	}
	defer Close()

	p0 := modules.
}

func TerminalSize() (int, int) {
	termbox.Sync()
	width, height := termbox.Size()
	return width, height
}

func Clear() {
	termbox.Clear(termbox.ColorDefault, termbox.Attribute(style.Theme.Default.Bg+1))
}

func Init() error {
	if err := termbox.Init(); err != nil {
		return err
	}
	termbox.SetInputMode(termbox.InputEsc | termbox.InputMouse)
	termbox.SetOutputMode(termbox.Output256)
	return nil
}

func Close() {
	termbox.Close()
}