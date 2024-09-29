package bootstrap

import (
	"github.com/LoveSnowEx/gotool/database/gormtool"
	"github.com/LoveSnowEx/gungi/config"
	"github.com/LoveSnowEx/gungi/internal/infra/database"
	"github.com/LoveSnowEx/gungi/internal/infra/po"
	"github.com/LoveSnowEx/gungi/tool/logger"
	"gorm.io/gorm"
)

func SetupDB() {
	db, err := gormtool.Open(config.Database, &gorm.Config{
		Logger: logger.Gorm(),
	})
	if err != nil {
		panic(err)
	}
	if err = db.AutoMigrate(
		&po.User{},
		&po.Game{},
		&po.Player{},
		&po.Piece{},
	); err != nil {
		panic(err)
	}
	database.SetDefault(db)
}
