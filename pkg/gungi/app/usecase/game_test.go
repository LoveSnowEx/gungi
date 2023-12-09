package usecase

import (
	"reflect"
	"testing"

	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"
)

func Test_gameUseCase_CreateGame(t *testing.T) {
	tests := []struct {
		name     string
		wantGame model.Game
		wantErr  bool
	}{
		{
			name:     "CreateGame",
			wantGame: model.NewGame(),
			wantErr:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewGameUseCase(nil)
			gotGame, err := u.CreateGame()
			if (err != nil) != tt.wantErr {
				t.Errorf("gameService.CreateGame() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotGame, tt.wantGame) {
				t.Errorf("gameService.CreateGame() = %v, want %v", gotGame, tt.wantGame)
			}
		})
	}
}

func Test_gameUseCase_StartGame(t *testing.T) {
	tests := []struct {
		name    string
		players []model.Player
		colors  []model.Color
		wantErr bool
	}{
		{
			name: "StartGame",
			players: []model.Player{
				model.NewPlayer("player1"),
				model.NewPlayer("player2"),
			},
			colors: []model.Color{
				model.White,
				model.Black,
			},
			wantErr: false,
		},
		{
			name:    "StartGameWithInvalidPlayers",
			players: []model.Player{},
			colors:  []model.Color{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := NewGameUseCase(nil)
			game, err := u.CreateGame()
			if err != nil {
				t.Errorf("gameService.CreateGame() error = %v", err)
				return
			}
			for i, player := range tt.players {
				if err := u.JoinGame(game.Id(), player.Id(), tt.colors[i]); err != nil {
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
