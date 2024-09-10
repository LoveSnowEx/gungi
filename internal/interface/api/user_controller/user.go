package user_controller

import "github.com/LoveSnowEx/gungi/internal/app/user_usecase"

type Controller struct {
	userUsecase *user_usecase.Usecase
}

func New(userUsecase *user_usecase.Usecase) *Controller {
	return &Controller{
		userUsecase: userUsecase,
	}
}
