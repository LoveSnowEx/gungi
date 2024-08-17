package repo

import "github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"

type GameRepo interface {
	// Find finds the game with the given id.
	Find(id uint) (game model.Game, err error)
	// Save saves the game.
	Save(game model.Game) (err error)
}
