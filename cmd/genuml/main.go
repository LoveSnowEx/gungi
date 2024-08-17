package main

import (
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	target := filepath.Join(wd, "uml/game.plantuml")
	f, err := os.Create(target)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	cmd := exec.Command("goplantuml", "-recursive", "-aggregate-private-members", "-show-aggregations", wd)
	cmd.Stdout = f
	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
