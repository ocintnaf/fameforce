package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ocintnaf/fameforce/config"
	"github.com/ocintnaf/fameforce/pkg/database"
	"github.com/ocintnaf/fameforce/repositories"
)

func main() {
	config, err := config.Init()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.NewConnection(config.Database)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	server := fiber.New()

	influencersRepository := repositories.NewInfluencerRepository(db)

	server.Get("/", func(ctx *fiber.Ctx) error {
		influencers, err := influencersRepository.FindAll()
		if err != nil {
			log.Fatal(err)
		}

		return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
			"data": influencers,
		})
	})

	server.Listen(":3000")
}
