package database

import (
	"gorm.io/gorm"
)

var (
	connections = make(map[string]*gorm.DB)
)

func SetDefault(d *gorm.DB) {
	connections["default"] = d
}

func Default() *gorm.DB {
	return connections["default"]
}

func Set(name string, d *gorm.DB) {
	connections[name] = d
}

func Get(name string) *gorm.DB {
	return connections[name]
}
