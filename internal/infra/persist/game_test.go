//go:build ignore

package persist_test

import (
	"log/slog"
	"runtime/debug"
	"testing"

	"github.com/LoveSnowEx/gungi/config"
	"github.com/LoveSnowEx/gungi/internal/bootstrap"
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_model"
	"github.com/LoveSnowEx/gungi/internal/infra/database"
	"github.com/LoveSnowEx/gungi/internal/infra/persist"
	"github.com/LoveSnowEx/gungi/internal/infra/po"
)

func setup() {
	config.Database.Database = ":memory:"
	bootstrap.SetupSlog()
	bootstrap.SetupDB()
}

func teardown() {
	sqlDB, _ := database.Default().DB()
	sqlDB.Close()
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

func Test_gameRepoImpl_Save(t *testing.T) {
	type args struct {
		game gungi_model.Game
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				game: gungi_model.NewGame(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := persist.NewGameRepo()
			if err := r.Save(tt.args.game); (err != nil) != tt.wantErr {
				t.Errorf("gameRepoImpl.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			gamePo := po.Game{}
			var gameCount int64
			if err := database.Default().Model(&gamePo).Count(&gameCount).Error; (err != nil) != tt.wantErr {
				t.Errorf("gameRepoImpl.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gameCount != 1 {
				t.Errorf("gameRepoImpl.Save() game count = %d, want 1", gameCount)
			}
		})
	}
}

func Test_gameRepoImpl_Save_Twice(t *testing.T) {
	type args struct {
		game gungi_model.Game
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				game: gungi_model.NewGame(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := persist.NewGameRepo()
			if err := r.Save(tt.args.game); (err != nil) != tt.wantErr {
				t.Errorf("gameRepoImpl.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			var gameCount int64
			if err := database.Default().Model(&po.Game{}).Count(&gameCount).Error; (err != nil) != tt.wantErr {
				t.Errorf("gameRepoImpl.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gameCount != 1 {
				t.Errorf("gameRepoImpl.Save() game count = %d, want 1", gameCount)
			}
			if tt.args.game.Id() == 0 {
				t.Errorf("game id is wrong, expected not 0, got %v", tt.args.game.Id())
			}
			if err := r.Save(tt.args.game); (err != nil) != tt.wantErr {
				t.Errorf("gameRepoImpl.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := database.Default().Model(&po.Game{}).Count(&gameCount).Error; (err != nil) != tt.wantErr {
				t.Errorf("gameRepoImpl.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gameCount != 1 {
				t.Errorf("gameRepoImpl.Save() game count = %d, want 1", gameCount)
			}
		})
	}
}

func Test_gameRepoImpl_Save_WithPlayer(t *testing.T) {
	type args struct {
		players []struct {
			name  string
			color gungi_model.Color
		}
		game gungi_model.Game
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				players: []struct {
					name  string
					color gungi_model.Color
				}{
					{
						name:  "test player",
						color: gungi_model.White,
					},
				},
				game: gungi_model.NewGame(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			database.Default().Exec()
			player := gungi_model.NewPlayer("test player")
			tt.args.game.Join(gungi_model.White, player)
			r := persist.NewGameRepo()
			if err := r.Save(tt.args.game); (err != nil) != tt.wantErr {
				t.Errorf("gameRepoImpl.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			var gameCount int64
			if err := database.Default().Model(&po.Game{}).Count(&gameCount).Error; (err != nil) != tt.wantErr {
				t.Errorf("gameRepoImpl.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gameCount != 1 {
				t.Errorf("gameRepoImpl.Save() game count = %d, want 1", gameCount)
			}
			if tt.args.game.Id() == 0 {
				t.Errorf("game id is wrong, expected not 0, got %v", tt.args.game.Id())
			}

			if err := r.Save(tt.args.game); (err != nil) != tt.wantErr {
				t.Errorf("gameRepoImpl.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := database.Default().Model(&po.Game{}).Count(&gameCount).Error; (err != nil) != tt.wantErr {
				t.Errorf("gameRepoImpl.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
			if gameCount != 1 {
				t.Errorf("gameRepoImpl.Save() game count = %d, want 1", gameCount)
			}
		})
	}
}
