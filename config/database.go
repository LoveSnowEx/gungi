package config

import (
	"path/filepath"

	"github.com/LoveSnowEx/gungi/tool/pathtool"
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
