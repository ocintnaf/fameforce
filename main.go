package main

import "github.com/gofiber/fiber/v2"

func main() {
	server := fiber.New()

	server.Get("/", func (ctx *fiber.Ctx) error {
		return ctx.Status(fiber.StatusOK).SendString("Ciao")
	})

	server.Listen(":3000")
}