package config

import "github.com/spf13/viper"

var (
	Database DatabaseConfig
)

type DatabaseConfig struct {
	Driver     string
	DataSource string
}

func init() {
	viper.SetDefault("database.driver", "sqlite3")
	viper.SetDefault("database.dataSource", "gungi.sqlite3")
	viper.AutomaticEnv()
	Database.Driver = viper.GetString("database.driver")
	Database.DataSource = viper.GetString("database.dataSource")
}
