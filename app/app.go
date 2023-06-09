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
	"github.com/ocintnaf/fameforce/pkg/logger"
	"github.com/ocintnaf/fameforce/pkg/middlewares"
	"github.com/ocintnaf/fameforce/repositories"
	"github.com/ocintnaf/fameforce/usecases"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type app struct {
	cfg      *config.Config
	fiber    *fiber.App
	db       *gorm.DB
	logger   *zap.Logger
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

func (a *app) initLogger() error {
	logger, err := logger.NewLogger(a.cfg.AppEnv)
	if err != nil {
		return err
	}

	a.logger = logger

	return nil
}

func (a *app) initDB() error {
	db, err := database.NewConnection(a.cfg.Database)
	if err != nil {
		return err
	}

	a.db = db

	return nil
}

func (a *app) initFirebaseApp() error {
	app, err := gcloud.NewFirebaseApp(a.cfg.GCloud)
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

func (a *app) initServer() error {
	a.fiber = fiber.New()

	a.registerRoutes()

	return a.fiber.Listen(":8080") // TODO: make port configurable
}

func (a *app) registerRoutes() {
	api := a.fiber.Group("/api")
	v1 := api.Group("/v1")

	authMiddleware := middlewares.NewAuthMiddleware(a.firebase.auth)

	// Health routes
	healthGroup := a.fiber.Group("/health")
	healthController := controllers.NewHealthController(a.fiber)
	healthGroup.Get("/", healthController.Health)

	// User routes
	userGroup := v1.Group("/users")
	userRepository := repositories.NewUserRepository(a.db)
	userUsecase := usecases.NewUserUsecase(userRepository)
	userController := controllers.NewUserController(userUsecase)

	userGroup.Get("/me", authMiddleware, userController.Me)
	userGroup.Post("/", authMiddleware, userController.Create)
}

func (a *app) Run() error {
	if err := a.initLogger(); err != nil {
		return err
	}

	if err := a.initDB(); err != nil {
		return err
	}

	if err := a.initFirebaseApp(); err != nil {
		return err
	}

	if err := a.initServer(); err != nil {
		return err
	}

	return nil
}

func (a *app) Shutdown() error {
	err := database.CloseConnection(a.db)
	if err != nil {
		return err
	}

	err = a.logger.Sync()
	if err != nil {
		return err
	}

	return nil
}
