package main

import (
	"log"

	"github.com/ocintnaf/fameforce/app"
	"github.com/ocintnaf/fameforce/config"
)

func main() {
	config, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	app := app.Init(config)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := app.Shutdown(); err != nil {
			log.Fatal(err)
		}
	}()
}
