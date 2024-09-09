package api

import (
	"github.com/LoveSnowEx/gungi/internal/app/gungi_usecase"
	"github.com/gofiber/fiber/v3"
)

func NewApp(gungi_usecase *gungi_usecase.GameUsecase) *fiber.App {
	app := fiber.New()
	return app
}
