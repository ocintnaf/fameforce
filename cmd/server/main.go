package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ocintnaf/fameforce/config"
	"github.com/ocintnaf/fameforce/controllers"
	"github.com/ocintnaf/fameforce/pkg/database"
	"github.com/ocintnaf/fameforce/repositories"
	"github.com/ocintnaf/fameforce/usecases"
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

	influencerGroup := v1.Group("/influencers")
	influencerRepository := repositories.NewInfluencerRepository(db)
	influencerUsecase := usecases.NewInfluencerUsecase(influencerRepository)
	influencerController := controllers.NewInfluencerController(influencerGroup, influencerUsecase)
	influencerController.RegisterRoutes()

	healthController := controllers.NewHealthController(v1)
	healthController.RegisterRoutes()

	server.Listen(":8080")
}
