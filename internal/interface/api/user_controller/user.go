package user_controller

import (
	"strconv"

	"github.com/LoveSnowEx/gungi/internal/app/user_usecase"
	"github.com/LoveSnowEx/gungi/internal/interface/api/user_dto"
	"github.com/gofiber/fiber/v3"
)

var _ Controller = (*controller)(nil)

type Config struct {
	UserUsecase user_usecase.Usecase
}

type Controller interface {
	Find(ctx fiber.Ctx) error
}

type controller struct {
	userUsecase user_usecase.Usecase
}

func New(config *Config) *controller {
	return &controller{
		userUsecase: config.UserUsecase,
	}
}

func (c *controller) Find(ctx fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return err
	}
	user, err := c.userUsecase.Find(uint(id))
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}
	return ctx.JSON(fiber.Map{
		"user": user_dto.User{
			Name: user.Name,
		},
	})
}
