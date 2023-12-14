package config

import (
	"path/filepath"

	"github.com/LoveSnowEx/gungi/tool"
	"github.com/spf13/viper"
)

var (
	Database DatabaseConfig
)

type DatabaseConfig interface {
	Driver() string
	SetDriver(string)
	Source() string
	SetSource(string)
}

type databaseConfig struct {
	driver string
	source string
}

func (c *databaseConfig) Driver() string {
	return c.driver
}

func (c *databaseConfig) Source() string {
	return c.source
}

func (c *databaseConfig) SetDriver(driver string) {
	c.driver = driver
}

func (c *databaseConfig) SetSource(source string) {
	c.source = filepath.Join(tool.ProjectRoot(), source)
}

func init() {
	viper.SetDefault("database.driver", "sqlite")
	viper.SetDefault("database.dataSource", "gungi.db")
	viper.AutomaticEnv()

	Database = &databaseConfig{}
	Database.SetDriver(viper.GetString("database.driver"))
	Database.SetSource(viper.GetString("database.dataSource"))
}
