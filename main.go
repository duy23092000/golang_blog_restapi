package main

import (
	"tutorial-rest/app"
	"tutorial-rest/config"
)

func main() {
	config := config.GetConfig()
	config.Port = ":8080"

	app := &app.App{}
	app.Initialize(config)
	app.Run(config.Port)
}
