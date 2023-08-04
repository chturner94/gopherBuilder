package main

import (
	"github.com/chturner94/gopherBuilder/cliutil/View"
	"time"
)

func main() {
	v := View.View{}
	v.Init()
	go v.Render()

	for {
		v.Canvas.Cell[0][0].Ch = 'X'
		time.Sleep(1 * time.Second)
	}
	//var c = Cliutil.UserControlState{}
	//go c.StartInputManager(Cliutil.MotionState, false)
	//
	//config := Config.Config{}
	//config.SetupConfig()
}
