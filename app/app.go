package app

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"github.com/gofiber/fiber/v2"
	"github.com/ocintnaf/fameforce/config"
	"github.com/ocintnaf/fameforce/controllers"
	"github.com/ocintnaf/fameforce/pkg/database"
	"github.com/ocintnaf/fameforce/pkg/gcloud"
	"github.com/ocintnaf/fameforce/pkg/middlewares"
	"github.com/ocintnaf/fameforce/repositories"
	"github.com/ocintnaf/fameforce/usecases"
	"gorm.io/gorm"
)

type app struct {
	cfg      *config.Config
	fiber    *fiber.App
	db       *gorm.DB
	firebase struct {
		app  *firebase.App
		auth *auth.Client
	}
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

	auth, err := app.Auth(context.Background())
	if err != nil {
		return err
	}

	a.firebase.app = app
	a.firebase.auth = auth

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

	authMiddleware := middlewares.NewAuthMiddleware(a.firebase.auth)

	// Health routes
	healthController := controllers.NewHealthController(a.fiber)
	healthController.RegisterRoutes()

	// User routes
	userGroup := v1.Group("/users")
	userRepository := repositories.NewUserRepository(a.db)
	userUsecase := usecases.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(userUsecase)

	userGroup.Get("/", authMiddleware, userController.GetAll)
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
