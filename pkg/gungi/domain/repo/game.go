package repo

import "github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"

type GameRepo interface {
	// Save saves the game.
	Save(game model.Game) (err error)
	// Find finds the game with the given ID.
	Find(id uint) (game model.Game, err error)
}
