package user_controller

import (
	"strconv"

	"github.com/LoveSnowEx/gungi/internal/app/user_usecase"
	"github.com/LoveSnowEx/gungi/internal/interface/api/user_dto"
	"github.com/gofiber/fiber/v3"
)

type Controller struct {
	userUsecase user_usecase.Usecase
}

func New(userUsecase user_usecase.Usecase) *Controller {
	return &Controller{
		userUsecase: userUsecase,
	}
}

func (c *Controller) Find(ctx fiber.Ctx) error {
	id, err := strconv.ParseUint(ctx.Params("id"), 10, 32)
	if err != nil {
		return err
	}
	user, err := c.userUsecase.Find(uint(id))
	if err != nil {
		ctx.Status(fiber.StatusNotFound)
		return ctx.JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	return ctx.JSON(fiber.Map{
		"user": user_dto.User{
			Name: user.Name(),
		},
	})
}
