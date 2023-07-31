package main

import (
	//"fmt"
	"github.com/chturner94/gopherBuilder/config"
	"github.com/chturner94/gopherBuilder/cliutil"
	//	"log"
	//	"encoding/json"
)

func main() {
	Cliutil.StartCli()
	config := Config.Config{}
	config.SetupConfig()
}
