package app

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/ocintnaf/fameforce/config"
	"github.com/ocintnaf/fameforce/controllers"
	"github.com/ocintnaf/fameforce/pkg/database"
	"github.com/ocintnaf/fameforce/repositories"
	"github.com/ocintnaf/fameforce/usecases"
	"gorm.io/gorm"
)

type app struct {
	cfg   *config.Config
	fiber *fiber.App
	ctx   context.Context
	db    *gorm.DB
}

func Init(cfg *config.Config) *app {
	return &app{
		cfg: cfg,
		ctx: context.Background(),
	}
}

func (a *app) initAndConnectDB() error {
	db, err := database.NewConnection(a.cfg.Database)
	if err != nil {
		return err
	}

	a.db = db

	return nil
}

func (a *app) initAndStartServer() error {
	a.fiber = fiber.New()

	a.registerRoutes()

	return a.fiber.Listen(":8080") // TODO: make port configurable
}

func (a *app) registerRoutes() {
	api := a.fiber.Group("/api")
	v1 := api.Group("/v1")

	// Health routes
	healthController := controllers.NewHealthController(a.fiber)
	healthController.RegisterRoutes()

	// Influencer routes
	influencerGroup := v1.Group("/influencers")
	influencerRepository := repositories.NewInfluencerRepository(a.db)
	influencerUsecase := usecases.NewInfluencerUsecase(influencerRepository)
	influencerController := controllers.NewInfluencerController(influencerGroup, influencerUsecase)
	influencerController.RegisterRoutes()
}

func (a *app) Run() error {
	if err := a.initAndConnectDB(); err != nil {
		return err
	}

	if err := a.initAndStartServer(); err != nil {
		return err
	}

	return nil
}

func (a *app) Shutdown() error {
	return database.CloseConnection(a.db)
}
