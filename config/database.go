package config

import (
	"path/filepath"

	"github.com/LoveSnowEx/gungi/tool/pathtool"
	"github.com/spf13/viper"
)

var (
	Database DatabaseConfig
)

type DatabaseConfig interface {
	Source() string
	SetSource(string)
}

type databaseConfig struct {
	source string
}

func (c *databaseConfig) Source() string {
	return c.source
}

func (c *databaseConfig) SetSource(source string) {
	c.source = filepath.Join(pathtool.ProjectRoot(), source)
}

func init() {
	viper.SetDefault("database.source", "gungi.db")
	viper.AutomaticEnv()

	Database = &databaseConfig{}
	Database.SetSource(viper.GetString("database.source"))
}
