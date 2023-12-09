package database

import (
	"github.com/LoveSnowEx/gungi/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func DB() *gorm.DB {
	if db != nil {
		return db
	}
	var err error
	db, err = gorm.Open(
		sqlite.Open(config.Database.DataSource),
	)
	if err != nil {
		panic(err)
	}
	return db
}
