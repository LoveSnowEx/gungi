package gungi_repo

import "github.com/LoveSnowEx/gungi/internal/domain/gungi_model"

type PlayerRepo interface {
	// Find finds the player with the given Id.
	Find(id uint) (player gungi_model.Player, err error)
}
