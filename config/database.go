package config

import "github.com/LoveSnowEx/gotool/database"

type DatabaseConfig struct {
	DefaultConnection string
	Connections       map[string]*database.Config
}
