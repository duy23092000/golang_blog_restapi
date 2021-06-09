package main

import (
	"tutorial-rest/app"
	"tutorial-rest/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(config.Port)
}
