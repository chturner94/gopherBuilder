package Cliutil

import (
	"fmt"
	//	"golang.org/x/term"
	"github.com/containerd/console"
)

const (
	MotionState = iota
	InputState
)

type KeyAction string

const (
	Up        KeyAction = "k"
	Down      KeyAction = "j"
	Left      KeyAction = "h"
	Right     KeyAction = "l"
	InputMode KeyAction = "i"
	Escape    KeyAction = "\033" // Escape key is represented by the escape character
)

var state int

func ReadUserInput() string {
	current := console.Current()
	defer current.Reset()
	if err := current.setRaw(); err != nil {
		panic(err)
	}
}

func DetectArrowKey(input KeyAction) KeyAction {
	switch state {
	case MotionState:
		switch input {
		case Up:
			return Up
		case Down:
			return Down
		case Right:
			return Right
		case Left:
			return Left
		case InputMode:
			state = InputState
			return InputMode
		default:
			return "Unknown"
		}
	case InputState:
		if input == Escape {
			state = MotionState
			return "Motion Mode"
		}
		// Process other input handling in the Input State if needed.
	}
	return "Unknown"
}

func StartCli() {
	state = MotionState
	fmt.Println("Welcome to Motion Mode (Use h(left) j(down) k(up) l(right) to move, i to enter Input Mode, and Esc to Move Again!)")
	for {
		userInput := ReadUserInput()
		action := DetectArrowKey(KeyAction(userInput))
		fmt.Println(action)
	}
}
