package main

import (
	"github.com/LoveSnowEx/gungi/internal/infra/database"
	"github.com/LoveSnowEx/gungi/internal/infra/persist"
	core "github.com/LoveSnowEx/gungi/pkg/core/app/usecase"
	gungi "github.com/LoveSnowEx/gungi/pkg/gungi/app/usecase"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"
)

func main() {
	database.Init()
	return
	coreApp := core.NewUserUsecase(&core.UserUsecaseConfig{
		UserRepo: persist.NewUserRepo(),
	})
	user1, err := coreApp.Create("test")
	if err != nil {
		panic(err)
	}
	user2, err := coreApp.Create("test2")
	if err != nil {
		panic(err)
	}
	gameApp := gungi.NewGameUsecase(&gungi.GameUsecaseConfig{
		GameRepo:   persist.NewGameRepo(),
		PlayerRepo: persist.NewPlayerRepo(),
	})
	game, err := gameApp.CreateGame()
	if err != nil {
		panic(err)
	}
	err = gameApp.JoinGame(game.Id(), user1.Id(), model.White)
	if err != nil {
		panic(err)
	}
	err = gameApp.JoinGame(game.Id(), user2.Id(), model.Black)
	if err != nil {
		panic(err)
	}
	err = gameApp.StartGame(game.Id())
	if err != nil {
		panic(err)
	}
}
