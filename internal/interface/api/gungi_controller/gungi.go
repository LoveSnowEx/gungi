package gungi_controller

import "github.com/LoveSnowEx/gungi/internal/app/gungi_usecase"

type Controller struct {
	gungiUsecase gungi_usecase.GameUsecase
}

func New(gungiUsecase gungi_usecase.GameUsecase) *Controller {
	return &Controller{
		gungiUsecase: gungiUsecase,
	}
}
