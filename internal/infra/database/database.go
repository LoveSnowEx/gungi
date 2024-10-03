package database

import (
	"github.com/uptrace/bun"
)

var (
	defaultConnection = ""
	connections       = make(map[string]*bun.DB)
)

func SetDefault(connection string) {
	defaultConnection = connection
}

func Default() *bun.DB {
	return connections[defaultConnection]
}

func Set(name string, d *bun.DB) {
	connections[name] = d
}

func Get(name string) *bun.DB {
	return connections[name]
}
