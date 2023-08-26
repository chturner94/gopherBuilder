package main

import (
	"github.com/chturner94/gopherBuilder/internal"
	"github.com/rivo/tview"
	"strconv"
)

type Application struct {
	internal.Configuration
}

func main() {
	app := Application{
		Configuration: *internal.InitConfig(),
	}
	var (
		init           = app.Initialized
		aws            = app.AwsEnabled
		github         = app.GithubEnabled
		serverName     = app.ServerHostName
		name           = app.FriendlyServerName
		path           = app.ConfigLocation
		ip             = app.ServerIP
		port           = app.ServicePort
		configLocation = app.ConfigLocation
	)
	newApp := tview.NewApplication()
	list := tview.NewList().
		AddItem("Initialized: "+strconv.FormatBool(init), "", 'a', nil).
		AddItem("AWS Enabled: "+strconv.FormatBool(aws), "", 'b', nil).
		AddItem("Github Enabled: "+strconv.FormatBool(github), "", 'c', nil).
		AddItem("Server Name: "+serverName, "", 'd', nil).
		AddItem("Friendly Server Name: "+name, "", 'e', nil).
		AddItem("Config Path: "+path, "", 'f', nil).
		AddItem("Server IP: "+ip, "", 'g', nil).
		AddItem("Server Port: "+strconv.Itoa(port), "", 'h', nil).
		AddItem("Config Location: "+configLocation, "", 'i', nil).
		AddItem("Quit", "Pres to exit", 'q', func() {
			newApp.Stop()
		})
	if err := newApp.SetRoot(list, true).SetFocus(list).Run(); err != nil {
		panic(err)
	}
}
