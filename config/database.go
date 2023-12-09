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
	viper.SetDefault("database.driver", "sqlite")
	viper.SetDefault("database.dataSource", "gungi.db")
	viper.AutomaticEnv()

	Database.Driver = viper.GetString("database.driver")
	Database.DataSource = viper.GetString("database.dataSource")
}
