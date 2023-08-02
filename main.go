package main

import (
	Cliutil "github.com/chturner94/gopherBuilder/cliutil"
	"github.com/chturner94/gopherBuilder/config"
	//	"log"
	//	"encoding/json"
)

func main() {

	var c = Cliutil.UserControlState{}
	go c.StartInputManager(Cliutil.MotionState, false)

	config := Config.Config{}
	config.SetupConfig()
}
