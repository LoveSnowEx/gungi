package main

import (
	"github.com/LoveSnowEx/gungi/internal/bootstrap"
	"github.com/LoveSnowEx/gungi/internal/interface/gui"
)

func main() {
	bootstrap.SetupSlog()
	bootstrap.SetupDB()

	if err := gui.Run(); err != nil {
		panic(err)
	}
}
