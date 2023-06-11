package controllers

import "github.com/gofiber/fiber/v2"

// BaseController defines the common properties of all fiber controllers.
type BaseController struct {
	Router fiber.Router
}
