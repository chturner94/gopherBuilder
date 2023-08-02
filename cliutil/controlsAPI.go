package Cliutil

import (
	"fmt"
	"sync"
)

// StartInputManager is the main function for the cliutil package for defining controls and managing user input.
// The StartInputManager function takes two parameters: initialState and locked. The initialState parameter is
// used to set the default mode of the user input. The locked parameter is used to lock the user input
// to the current mode. This is useful for when you want to lock the user input to a specific mode
// during a specific operation.
func (s *UserControlState) StartInputManager(initialState keybindMapping, locked bool) {
	s.Lock()
	defer s.Unlock()
	s.mode = initialState
	s.locked = locked

	fmt.Println("Welcome to Motion Mode (Use h(left) j(down) k(up) l(right) to move, i to enter Input Mode, and Esc to Move Again!)")
	for {
		userInput := readUserInput()
		action := detectInput(inputKey(userInput), s)
		fmt.Println(action)
	}
}

// UserControlState is the struct that is used to manage the state of the user input. It is used to
// manage the current mode of the user input, and to lock the user input to a specific mode.
// UserControlState needs to be declared in the main function of your program, and you can then run
// the StartInputManager function on the declared UserControlState variable. You can also run the LockModeSwitching
// and UnlockModeSwitching functions on the declared UserControlState variable to lock and unlock the
// user input to the current mode.
type UserControlState struct {
	sync.Mutex
	mode   keybindMapping
	locked bool
}

// LockModeSwitching is used to lock the user input to the current mode. This is useful for when you want
// to lock the user input to a specific mode during a specific operation.
func (s *UserControlState) LockModeSwitching() {
	s.Lock()
	defer s.Unlock()
	s.locked = true
}

// UnlockModeSwitching is used to unlock the user input to the current mode. This is useful for when you want
// to unlock the user input.
func (s *UserControlState) UnlockModeSwitching() {
	s.Lock()
	defer s.Unlock()
	s.locked = false
}
