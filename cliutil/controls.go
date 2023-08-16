package cliutil

import (
	"bufio"
	"fmt"
	"golang.org/x/term"
	"os"
	"os/signal"
	"syscall"
)

// The keybindMapping type is an enum that represents the current input mode of the user.
// This works similarly to vim's modes, where the user can either be in motion mode or input mode.
// To set the default mode of StartCli, use the MotionState or InputState constants.
type keybindMapping int

const (
	MotionState keybindMapping = iota
	InputState
)

// The inputKey type allows you to define input characters that are made available outputMap for remapping
// keys. These mappings are only enabled during motion mode, and otherwise will work normally in input mode.
type inputKey string

const (
	Up         inputKey = "k"
	Down       inputKey = "j"
	Left       inputKey = "h"
	Right      inputKey = "l"
	InputMode  inputKey = "i"
	MotionMode inputKey = "\033" // Escape key is represented by the escape character
)

// The outputMap type is the expected type for the keymap variable, which is used for remapping keys.
type outputMap string

// The keymap variable expects the inputKey type as a key, and a string value as the outputMap type. This
// is later used in the detectInput function to remap keys.
var keymap = map[inputKey]outputMap{
	Up:         "\033[1A",
	Down:       "\033[1B",
	Right:      "\033[1C",
	Left:       "\033[1D",
	InputMode:  "",
	MotionMode: "",
}

// The passthroughsInputKey function is used when we wish to override the outputMap for a given inputKey,
// or to pass through any key that is not defined in the keymap variable.
func passthroughsInputKey(input inputKey) outputMap {
	return outputMap(input)
}

// The readUserInput function is used to read a single character from the terminal. It also sets the
// terminal to raw mode, which allows us to read a single character at a time.
func readUserInput() string {
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

// The detectInput function is used to determine what action to take based on the inputKey type. If the
// inputKey type is defined in the keymap variable, then the outputMap type is returned. Otherwise, the
// inputKey type is passed through to the passthroughsInputKey function.
func detectInput(input inputKey, stateObject *UserControlState) outputMap {

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
		return passthroughsInputKey(input)
	case InputState:
		if !stateObject.locked && input == MotionMode {
			stateObject.mode = MotionState
			return "Motion Mode"
		}
		fmt.Print("Input State:", input)
		return outputMap(input)
	}
	return ""
}
