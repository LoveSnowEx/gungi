package service

import (
	"reflect"
	"testing"

	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/repo"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/service"
	"github.com/LoveSnowEx/gungi/pkg/gungi/infra/persist"
)

func Test_gameService_CreateGame(t *testing.T) {
	type fields struct {
		gameRepo    repo.GameRepo
		gameService service.GameService
		playerRepo  repo.PlayerRepo
	}
	tests := []struct {
		name     string
		fields   fields
		wantGame model.Game
		wantErr  bool
	}{
		{
			name: "CreateGame",
			fields: fields{
				gameRepo:    persist.NewGameRepo(),
				gameService: service.NewGameService(),
				playerRepo:  persist.NewPlayerRepo(),
			},
			wantGame: model.NewGame(),
			wantErr:  false,
		},
		{
			name: "InvalidServiceCreateGame",
			fields: fields{
				gameRepo:    nil,
				gameService: nil,
				playerRepo:  nil,
			},
			wantGame: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &gameService{
				gameRepo:    tt.fields.gameRepo,
				gameService: tt.fields.gameService,
				playerRepo:  tt.fields.playerRepo,
			}
			gotGame, err := c.CreateGame()
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

func Test_gameService_StartGame(t *testing.T) {
	type fields struct {
		gameRepo    repo.GameRepo
		gameService service.GameService
		playerRepo  repo.PlayerRepo
	}
	type args struct {
		gameId uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "StartGame",
			fields: fields{
				gameRepo:    persist.NewGameRepo(),
				gameService: service.NewGameService(),
				playerRepo:  persist.NewPlayerRepo(),
			},
			args: args{
				gameId: 1,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &gameService{
				gameRepo:    tt.fields.gameRepo,
				gameService: tt.fields.gameService,
				playerRepo:  tt.fields.playerRepo,
			}
			if err := c.StartGame(tt.args.gameId); (err != nil) != tt.wantErr {
				t.Errorf("gameService.StartGame() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
