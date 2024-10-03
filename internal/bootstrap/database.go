package bootstrap

import (
	"database/sql"

	"github.com/LoveSnowEx/gungi/config"
	"github.com/LoveSnowEx/gungi/internal/infra/database"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func SetupDB() {
	for connection, config := range config.Database.Connections {
		switch config.Driver {
		case "sqlite":
			sqldb, err := sql.Open(sqliteshim.ShimName, config.Dsn())
			if err != nil {
				panic(err)
			}
			database.Set(connection, bun.NewDB(sqldb, sqlitedialect.New()))
		}
	}
	database.SetDefault(config.Database.DefaultConnection)
}
