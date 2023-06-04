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

	a := app.Init(config)

	if err := a.Run(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := a.Shutdown(); err != nil {
			log.Fatal(err)
		}
	}()
}
