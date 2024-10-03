package config

import (
	"io"
	"os"
	"path/filepath"

	"github.com/LoveSnowEx/gungi/tool/pathtool"
	"github.com/spf13/viper"
)

var (
	Database DatabaseConfig
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
	viper.SetDefault("database.default", "")
	Database.DefaultConnection = viper.GetString("database.default")
	if err := viper.Sub("database.connections").Unmarshal(&Database.Connections); err != nil {
		panic(err)
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
