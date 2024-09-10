package user_controller

import "github.com/LoveSnowEx/gungi/internal/app/user_usecase"

type Controller struct {
	userUsecase *user_usecase.UserUsecase
}

func New(userUsecase *user_usecase.UserUsecase) *Controller {
	return &Controller{
		userUsecase: userUsecase,
	}
}
