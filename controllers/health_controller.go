package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type healthController struct {
	BaseController
}

// HealthController is the interface for the health controller.
type HealthController interface {
	Health(ctx *fiber.Ctx) error
}

// NewHealthController returns a new health controller.
func NewHealthController(router fiber.Router) *healthController {
	return &healthController{
		BaseController: BaseController{
			Router: router,
		},
	}
}

// Health checks the health status of the application.
// Returns a 200 OK status code if the application is healthy.
// Return a 503 Service Unavailable status code if the application is unhealthy.
func (hc *healthController) Health(ctx *fiber.Ctx) error {
	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":    "ok",
		"timestamp": time.Now(),
	})
}

// RegisterRoutes registers the health controller routes.
func (hc *healthController) RegisterRoutes() {
	hc.Router.Get("/health", hc.Health)
}
