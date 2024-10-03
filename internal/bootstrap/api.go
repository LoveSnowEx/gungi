package bootstrap

import (
	"github.com/LoveSnowEx/gungi/internal/app/gungi_usecase"
	"github.com/LoveSnowEx/gungi/internal/app/user_usecase"
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_service"
	"github.com/LoveSnowEx/gungi/internal/infra/database"
	"github.com/LoveSnowEx/gungi/internal/infra/notification"
	"github.com/LoveSnowEx/gungi/internal/infra/persist"
	"github.com/LoveSnowEx/gungi/internal/interface/api/gungi_controller"
	"github.com/LoveSnowEx/gungi/internal/interface/api/route"
	"github.com/LoveSnowEx/gungi/internal/interface/api/user_controller"
	"github.com/gofiber/fiber/v3"
)

func SetupApi(app *fiber.App) {
	route.Setup(app, &route.Config{
		UserController: user_controller.New(&user_controller.Config{
			UserUsecase: user_usecase.New(&user_usecase.Config{
				UserRepo: persist.NewUserRepo(database.Default()),
			}),
		}),
		GungiController: *gungi_controller.New(&gungi_controller.Config{
			GungiUsecase: gungi_usecase.New(&gungi_usecase.Config{
				GameService: gungi_service.NewGameService(),
				// GameRepo:     persist.NewGameRepo(),
				EventManager: notification.NewGameManager(),
			}),
		}),
	})
}
