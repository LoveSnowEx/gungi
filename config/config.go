package config

import (
	"io"
	"os"
	"path/filepath"

	"github.com/LoveSnowEx/gotool/database"
	"github.com/LoveSnowEx/gungi/tool/pathtool"
	"github.com/spf13/viper"
)

var (
	Database *database.Config
	Logger   LoggerConfig
)

func init() {
	configPath := filepath.Join(pathtool.ProjectRoot(), "config")
	// viper.AutomaticEnv()
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// database
	viper.SetDefault("database.driver", "sqlite")
	viper.SetDefault("database.database", "tmp/gungi.db")
	if config, err := database.ReadViper(viper.Sub("database")); err != nil {
		panic(err)
	} else {
		Database = config
	}

	// logger
	viper.SetDefault("logger.slog", nil)
	viper.SetDefault("logger.gorm", nil)
	Logger = &loggerConfig{
		wrtiers: map[string]io.Writer{
			"slog": os.Stdout,
			"gorm": os.Stdout,
		},
	}
	loadLoggerConfig(Logger)
}
