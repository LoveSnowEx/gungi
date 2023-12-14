package bootstrap

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/LoveSnowEx/gungi/config"
	"github.com/LoveSnowEx/gungi/internal/infra/dal"
	"github.com/LoveSnowEx/gungi/internal/infra/database"
	"github.com/LoveSnowEx/gungi/internal/infra/po"
	"github.com/LoveSnowEx/gungi/tool/logger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupSlog() {
	slog.SetDefault(logger.Slog())
}

func SetupDB() {
	source := config.Database.Source()
	dir := filepath.Dir(source)
	var err error
	if _, err = os.Stat(dir); err != nil {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			panic(err)
		}
	}

	db, err := gorm.Open(
		sqlite.Open(source),
		&gorm.Config{
			Logger: logger.Gorm(),
		},
	)
	if err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(
		&po.User{},
		&po.Game{},
		&po.Player{},
		&po.Piece{},
		&po.BoardPiece{},
		&po.ReservePiece{},
		&po.DiscardPiece{},
	); err != nil {
		panic(err)
	}
	database.SetDefault(db)
	dal.SetDefault(db)
}
