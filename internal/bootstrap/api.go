package bootstrap

import (
	"github.com/LoveSnowEx/gungi/internal/interface/api/route"
	"github.com/gofiber/fiber/v3"
)

func SetupApi(app *fiber.App) {
	route.Setup(app, nil)
}
