package app

import (
	"context"
	"database/sql"

	"github.com/gofiber/fiber/v2"
	"github.com/ocintnaf/fameforce/config"
	"github.com/ocintnaf/fameforce/controllers"
	"github.com/ocintnaf/fameforce/pkg/database"
	"github.com/ocintnaf/fameforce/repositories"
	"github.com/ocintnaf/fameforce/usecases"
)

type app struct {
	cfg   *config.Config
	fiber *fiber.App
	ctx   context.Context
	db    *sql.DB
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

	return a.fiber.Listen(":8080") // TODO: make port configurable
}

func (a *app) RegisterRoutes() {
	api := a.fiber.Group("/api")
	v1 := api.Group("/v1")

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
	return a.db.Close()
}
