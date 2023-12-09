package main

import (
	"github.com/LoveSnowEx/gungi/internal/infra/database"
	"github.com/LoveSnowEx/gungi/pkg/gungi/app/usecase"
)

func main() {
	db := database.DB()
	defer db.Close()
	app := usecase.NewGameUseCase(&usecase.GameUseCaseConfig{})
	app.CreateGame()
}
