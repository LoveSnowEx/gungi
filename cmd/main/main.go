package main

import (
	"log"

	"github.com/LoveSnowEx/gungi/internal/bootstrap"
	"github.com/gofiber/fiber/v3"
)

func main() {
	bootstrap.SetupSlog()
	bootstrap.SetupDB()

	app := fiber.New()
	bootstrap.SetupApi(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalln(err)
	}
}
