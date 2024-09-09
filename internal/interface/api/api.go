package api

import "github.com/gofiber/fiber/v3"

func NewApp() *fiber.App {
	app := fiber.New()
	return app
}
