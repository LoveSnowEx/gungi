package usecase

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/LoveSnowEx/gungi/config"
	"github.com/LoveSnowEx/gungi/internal/infra/database"
	"github.com/LoveSnowEx/gungi/internal/infra/persist"
	coremodel "github.com/LoveSnowEx/gungi/pkg/core/domain/model"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"
	"github.com/google/uuid"
)

func setup() {
	u, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	source := filepath.Join("tmp", u.String()+".db")
	config.Database.SetSource(source)
	database.Init()
}

func teardown() {
	database.Close()
	_ = os.Remove(config.Database.Source())
}

func TestMain(m *testing.M) {
	defer func() {
		err := recover()
		if err != nil {
			teardown()
			panic(err)
		}
	}()
	setup()
	m.Run()
	teardown()
}

func NewFakeGameUsecase() GameUsecase {
	return NewGameUsecase(&GameUsecaseConfig{
		GameRepo:   persist.NewGameRepo(),
		PlayerRepo: persist.NewPlayerRepo(),
	})
}

func MakeFakeUsers(users ...coremodel.User) (err error) {
	for _, user := range users {
		if err = persist.NewUserRepo().Create(user); err != nil {
			return
		}
	}
	return
}

func Test_gameUsecase_CreateGame(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "CreateGame",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewFakeGameUsecase()
			gotGame, err := u.CreateGame()
			if err != nil {
				t.Errorf("gameService.CreateGame() got error = %v", err)
				return
			}
			if gotGame.Id() == 0 {
				t.Errorf("game id is wrong, got %v", gotGame.Id())
			}
		})
	}
}

func Test_gameUsecase_StartGame(t *testing.T) {
	tests := []struct {
		name        string
		playernames []string
		colors      []model.Color
		wantErr     bool
	}{
		{
			name: "StartGame",
			playernames: []string{
				"test player 1",
				"test player 2",
			},
			colors: []model.Color{
				model.White,
				model.Black,
			},
			wantErr: false,
		},
		{
			name:        "StartGameWithInvalidPlayers",
			playernames: []string{},
			colors:      []model.Color{},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewFakeGameUsecase()
			game, err := u.CreateGame()
			if err != nil {
				t.Errorf("gameService.CreateGame() error = %v", err)
				return
			}
			for i, playername := range tt.playernames {
				user := coremodel.NewUser(playername)
				err = MakeFakeUsers(user)
				if err != nil {
					t.Errorf("MakeFakeUsers() error = %v", err)
					return
				}
				if err := u.JoinGame(game.Id(), user.Id(), tt.colors[i]); err != nil {
					t.Errorf("gameService.JoinGame() error = %v", err)
					return
				}
			}
			if err := u.StartGame(game.Id()); (err != nil) != tt.wantErr {
				t.Errorf("gameService.StartGame() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
