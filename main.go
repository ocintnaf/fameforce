package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ocintnaf/fameforce/config"
	"github.com/ocintnaf/fameforce/pkg/database"
)

func main() {
	config, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	log.Print(config)

	db, err := database.NewConnection(config.Database)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	server := fiber.New()

	server.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("Ciao")
	})

	server.Listen(":3000")
}
