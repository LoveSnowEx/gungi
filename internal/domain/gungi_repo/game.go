package gungi_repo

import "github.com/LoveSnowEx/gungi/internal/domain/gungi_model"

type GameRepo interface {
	// Find finds the game with the given id.
	Find(id uint) (game gungi_model.Game, err error)
	// Save saves the game.
	Save(game gungi_model.Game) (err error)
}
