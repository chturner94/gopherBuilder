package Cliutil

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type KeybindMapping int

const (
	MotionState KeybindMapping = iota
	InputState
)

type UserControlState struct {
	sync.Mutex
	mode   KeybindMapping
	locked bool
}

//var state = UserControlState{mode: MotionState}

func (s *UserControlState) LockModeSwitching() {
	s.Lock()
	defer s.Unlock()
	s.locked = true
}

func (s *UserControlState) UnlockModeSwitching() {
	s.Lock()
	defer s.Unlock()
	s.locked = false
}

//var (
//	state      int
//	stateMutex sync.Mutex
//	stateLocked bool
//)

type InputKey string

const (
	Up         InputKey = "k"
	Down       InputKey = "j"
	Left       InputKey = "h"
	Right      InputKey = "l"
	InputMode  InputKey = "i"
	MotionMode InputKey = "\033" // Escape key is represented by the escape character
)

type OutputMap string

var keymap = map[InputKey]OutputMap{
	Up:         "\033[1A",
	Down:       "\033[1B",
	Right:      "\033[1C",
	Left:       "\033[1D",
	InputMode:  "",
	MotionMode: "",
}

func handleUndefinedKey(input InputKey) OutputMap {
	return OutputMap(input)
}

//func LockState() {
//	stateMutex.Lock()
//	defer stateMutex.Unlock()
//	stateLocked = true
//}

// func UnlockState() {
// 	stateMutex.Lock()
// 	defer stateMutex.Unlock()
// 	stateLocked = false
// }

func ReadUserInput() string {
	// Get the terminal state to enable raw mode
	oldState, err := term.MakeRaw(0)
	if err != nil {
		fmt.Println("Error setting raw mode:", err)
		os.Exit(1)
	}

	// Restore the terminal state when the function exits
	defer term.Restore(0, oldState)

	// Setup signal channel to listen for interrupts (i.e. Ctrl+c or Ctrl+z)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigChan
		switch sig {
		case syscall.SIGINT:
			fmt.Println("\nReceived SIGINT (Ctrl+C). Exiting...")
			os.Exit(0)
		case syscall.SIGTERM:
			fmt.Println("\nReceived SIGTERM (Ctrl+Z). Exiting...")
			os.Exit(0)
		}
	}()

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	return string(char)
}

// func DetectInput(input InputKey) OutputMap {
// 	stateMutex.Lock()
// 	defer stateMutex.Unlock()

// 	switch state {
// 	case MotionState:
// 		if !stateLocked && input == InputMode {
// 			state = InputState
// 			return "Input Mode"
// 		}
// 		if mappedKey, ok := keymap[input]; ok {
// 			fmt.Print("Motion State:", mappedKey)
// 			return mappedKey
// 		}
// 		return handleUndefinedKey(input)
// 	case InputState:
// 		if !stateLocked && input == MotionMode {
// 			state = MotionState
// 			return "Motion Mode"
// 		}
// 		fmt.Print("Input State:", input)
// 		return OutputMap(input)
// 	}
// 	return ""
// }

func DetectInput(input InputKey, stateObject *UserControlState) OutputMap {

	switch stateObject.mode {
	case MotionState:
		if !stateObject.locked && input == InputMode {
			stateObject.mode = InputState
			return "Input Mode"
		}
		if mappedKey, ok := keymap[input]; ok {
			fmt.Print("Motion State:", mappedKey)
			return mappedKey
		}
		return handleUndefinedKey(input)
	case InputState:
		if !stateObject.locked && input == MotionMode {
			stateObject.mode = MotionState
			return "Motion Mode"
		}
		fmt.Print("Input State:", input)
		return OutputMap(input)
	}
	return ""
}

func (s *UserControlState) StartCli(initialState KeybindMapping, locked bool) {
	s.Lock()
	defer s.Unlock()
	s.mode = initialState
	s.locked = locked

	fmt.Println("Welcome to Motion Mode (Use h(left) j(down) k(up) l(right) to move, i to enter Input Mode, and Esc to Move Again!)")
	for {
		userInput := ReadUserInput()
		action := DetectInput(InputKey(userInput), s)
		fmt.Println(action)
	}
}
