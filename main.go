package main

import (
	//"fmt"
	"github.com/chturner94/gopherBuilder/cliutil"
	"github.com/chturner94/gopherBuilder/config"
	//	"log"
	//	"encoding/json"
)

func main() {
	var c = Cliutil.UserControlState{}
	c.StartCli(Cliutil.MotionState, false)

	config := Config.Config{}
	config.SetupConfig()
}
