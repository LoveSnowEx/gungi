package main

import (
	"github.com/LoveSnowEx/gungi/pkg/gungi/app/usecase"
)

func main() {
	app := usecase.NewGameUseCase(&usecase.GameUseCaseConfig{})
	app.CreateGame()
}
