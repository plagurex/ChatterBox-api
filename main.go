package main

import (
	"chatterbox/internal/app"
	"chatterbox/internal/utils"
	"log"
)

func main() {
	a := app.NewApp()
	log.Fatal(a.Run(utils.LoadConfig("config.json")))
}
