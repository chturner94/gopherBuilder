package main

import (
	Config "github.com/chturner94/gopherBuilder/config"
	//	"log"
	//	"encoding/json"
)

func main() {

	var c = Cliut.UserControlState{}
	go c.StartInputManager(Cliutil2.MotionState, false)

	config := Config.Config{}
	config.SetupConfig()
}
