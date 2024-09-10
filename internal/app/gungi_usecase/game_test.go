package gungi_usecase_test

import (
	"log/slog"
	"os"
	"path/filepath"
	"runtime/debug"
	"testing"

	. "github.com/LoveSnowEx/gungi/internal/app/gungi_usecase"

	"github.com/LoveSnowEx/gungi/config"
	"github.com/LoveSnowEx/gungi/internal/bootstrap"
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_model"
	"github.com/LoveSnowEx/gungi/internal/domain/user_model"
	"github.com/LoveSnowEx/gungi/internal/infra/database"
	"github.com/LoveSnowEx/gungi/internal/infra/notification"
	"github.com/LoveSnowEx/gungi/internal/infra/persist"
	"github.com/google/uuid"
	"github.com/gookit/event"
)

func setup() {
	u, err := uuid.NewUUID()
	if err != nil {
		panic(err)
	}
	source := filepath.Join("tmp", u.String()+".db")
	config.Database.SetSource(source)
	bootstrap.SetupSlog()
	bootstrap.SetupDB()
}

func teardown() {
	sqlDB, _ := database.Default().DB()
	sqlDB.Close()
	_ = os.Remove(config.Database.Source())
}

func TestMain(m *testing.M) {
	defer func() {
		if err := recover(); err != nil {
			slog.Debug("Panic!", "err", err, "stacktrace", string(debug.Stack()))
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
		GameRepo:     persist.NewGameRepo(),
		PlayerRepo:   persist.NewPlayerRepo(),
		EventManager: notification.NewGameManager(),
	})
}

func MakeFakeUsers(users ...user_model.User) (err error) {
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
		colors      []gungi_model.Color
		wantErr     bool
	}{
		{
			name: "StartGame",
			playernames: []string{
				"test player 1",
				"test player 2",
			},
			colors: []gungi_model.Color{
				gungi_model.White,
				gungi_model.Black,
			},
			wantErr: false,
		},
		{
			name:        "StartGameWithInvalidPlayers",
			playernames: []string{},
			colors:      []gungi_model.Color{},
			wantErr:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			updateCount := 0
			expectedUpdateCount := 0
			notification.NewGameManager().AddListener("game.update.*", event.ListenerFunc(func(_ event.Event) error {
				updateCount++
				return nil
			}))
			u := NewFakeGameUsecase()
			game, err := u.CreateGame()
			if err != nil {
				t.Errorf("gameService.CreateGame() error = %v", err)
				return
			}
			for i, playername := range tt.playernames {
				user := user_model.NewUser(playername)
				err = MakeFakeUsers(user)
				if err != nil {
					t.Errorf("MakeFakeUsers() error = %v", err)
					return
				}
				if err := u.JoinGame(game.Id(), user.Id(), tt.colors[i]); err != nil {
					t.Errorf("gameService.JoinGame() error = %v", err)
					return
				}
				expectedUpdateCount++
			}
			if err := u.StartGame(game.Id()); (err != nil) != tt.wantErr {
				t.Errorf("gameService.StartGame() error = %v, wantErr %v", err, tt.wantErr)
			}
			expectedUpdateCount++
			if updateCount != expectedUpdateCount && !tt.wantErr {
				t.Errorf("gameService should have fired %v events, but fired %v", expectedUpdateCount, updateCount)
			}
		})
	}
}
