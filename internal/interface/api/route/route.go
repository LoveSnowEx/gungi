package route

import (
	"github.com/LoveSnowEx/gungi/internal/interface/api/gungi_controller"
	"github.com/LoveSnowEx/gungi/internal/interface/api/user_controller"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/healthcheck"
)

type Config struct {
	UserController  user_controller.Controller
	GungiController gungi_controller.Controller
}

func Setup(app *fiber.App, config *Config) {
	// Health check
	app.Get(healthcheck.DefaultLivenessEndpoint, healthcheck.NewHealthChecker())

	// User
	userGroup := app.Group("/user")
	userGroup.Get("/:id", config.UserController.Find)
	userGroup.Post("/", config.UserController.Create)

	// TODO: Add routes here
}
