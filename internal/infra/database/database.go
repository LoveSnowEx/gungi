package database

import (
	"database/sql"

	"github.com/LoveSnowEx/gungi/config"
)

var (
	db *sql.DB
)

func DB() *sql.DB {
	if db != nil {
		return db
	}
	databaseConfig := config.Database
	var err error
	db, err = sql.Open(databaseConfig.Driver, databaseConfig.DataSource)
	if err != nil {
		panic(err)
	}
	return db
}
