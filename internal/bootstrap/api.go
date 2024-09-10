package bootstrap

import (
	"log/slog"

	"github.com/gofiber/fiber/v3/middleware/logger"

	"github.com/LoveSnowEx/gungi/internal/app/user_usecase"
	"github.com/LoveSnowEx/gungi/internal/infra/persist"
	"github.com/LoveSnowEx/gungi/internal/interface/api/route"
	"github.com/LoveSnowEx/gungi/internal/interface/api/user_controller"
	"github.com/gofiber/fiber/v3"
)

func SetupApi(app *fiber.App) {
	app.Use(logger.New(
		logger.ConfigDefault,
	))
	app.Use(func(c fiber.Ctx) (err error) {
		err = c.Next()
		if err != nil {
			slog.Default().Error("api error", "err", err.Error())
		}
		return
	})
	route.Setup(app, &route.Config{
		UserController: user_controller.New(
			user_usecase.New(&user_usecase.Config{
				UserRepo: persist.NewUserRepo(),
			}),
		),
	})
}
