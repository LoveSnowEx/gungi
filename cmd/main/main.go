package main

import (
	"github.com/LoveSnowEx/gungi/internal/bootstrap"
)

func main() {
	bootstrap.SetupSlog()
	bootstrap.SetupDB()
}
