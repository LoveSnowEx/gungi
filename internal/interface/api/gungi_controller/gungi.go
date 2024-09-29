package gungi_controller

import "github.com/LoveSnowEx/gungi/internal/app/gungi_usecase"

type Config struct {
	GungiUsecase gungi_usecase.GameUsecase
}

type Controller interface {
}

type controller struct {
	gungiUsecase gungi_usecase.GameUsecase
}

func New(config *Config) *controller {
	return &controller{
		gungiUsecase: config.GungiUsecase,
	}
}
