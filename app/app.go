package app

import (
	firebase "firebase.google.com/go/v4"
	"github.com/gofiber/fiber/v2"
	"github.com/ocintnaf/fameforce/config"
	"github.com/ocintnaf/fameforce/controllers"
	"github.com/ocintnaf/fameforce/pkg/database"
	"github.com/ocintnaf/fameforce/pkg/gcloud"
	"github.com/ocintnaf/fameforce/repositories"
	"github.com/ocintnaf/fameforce/usecases"
	"gorm.io/gorm"
)

type app struct {
	cfg      *config.Config
	fiber    *fiber.App
	db       *gorm.DB
	firebase *firebase.App
}

func Init(cfg *config.Config) *app {
	return &app{
		cfg: cfg,
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

func (a *app) initAndConnectFirebaseApp() error {
	app, err := gcloud.NewFirebaseApp()
	if err != nil {
		return err
	}

	a.firebase = app

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

	// User routes
	userGroup := v1.Group("/users")
	userRepository := repositories.NewUserRepository(a.db)
	userUsecase := usecases.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(userGroup, userUsecase)
	userController.RegisterRoutes()
}

func (a *app) Run() error {
	if err := a.initAndConnectDB(); err != nil {
		return err
	}

	if err := a.initAndConnectFirebaseApp(); err != nil {
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
