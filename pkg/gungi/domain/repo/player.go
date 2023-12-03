package repo

import "github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"

type PlayerRepo interface {
	// Save saves the player.
	Save(player model.Player) (err error)
	// Find finds the player with the given ID.
	Find(id uint) (player model.Player, err error)
}
