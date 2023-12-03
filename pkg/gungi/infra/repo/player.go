package repo

import (
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/repo"
	"github.com/LoveSnowEx/gungi/pkg/gungi/infra/errors"
)

var playerStorage = make(map[uint]model.Player)

type playerRepoImpl struct {
}

func NewPlayerRepo() repo.PlayerRepo {
	return &playerRepoImpl{}
}

func (r *playerRepoImpl) Save(player model.Player) (err error) {
	playerStorage[player.ID()] = player
	return
}

func (r *playerRepoImpl) Find(id uint) (player model.Player, err error) {
	player, ok := playerStorage[id]
	if !ok {
		err = errors.ErrPlayerNotFound
	}
	return player, err
}
