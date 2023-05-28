package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ocintnaf/fameforce/config"
	"github.com/ocintnaf/fameforce/controllers"
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

	api := server.Group("/api")

	v1 := api.Group("/v1")

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

	healthController := controllers.NewHealthController(v1)
	healthController.RegisterRoutes()

	server.Listen(":8080")
}
