package repo

import "github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"

type PlayerRepo interface {
	// Find finds the player with the given Id.
	Find(id uint) (player model.Player, err error)
}
