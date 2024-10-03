package database

import (
	"github.com/uptrace/bun"
)

var (
	connections = make(map[string]*bun.DB)
)

func SetDefault(d *bun.DB) {
	connections["default"] = d
}

func Default() *bun.DB {
	return connections["default"]
}

func Set(name string, d *bun.DB) {
	connections[name] = d
}

func Get(name string) *bun.DB {
	return connections[name]
}
