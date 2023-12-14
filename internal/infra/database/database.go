package database

import (
	"os"
	"path/filepath"

	"github.com/LoveSnowEx/gungi/config"
	"github.com/LoveSnowEx/gungi/internal/infra/dal"
	"github.com/LoveSnowEx/gungi/internal/infra/po"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

// Init database and migrate tables.
func Init() {
	var err error
	if db != nil {
		return
	}

	source := config.Database.Source()
	dir := filepath.Dir(source)
	if _, err = os.Stat(dir); err != nil {
		if err = os.MkdirAll(dir, os.ModePerm); err != nil {
			panic(err)
		}
	}

	if db, err = gorm.Open(sqlite.Open(config.Database.Source())); err != nil {
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
	dal.SetDefault(db)
}

// Close database.
func Close() {
	if db == nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}
	if err = sqlDB.Close(); err != nil {
		panic(err)
	}
	db = nil
}
