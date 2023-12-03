package repo

import (
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/repo"
	"github.com/LoveSnowEx/gungi/pkg/gungi/infra/errors"
)

var gameStorage = make(map[uint]model.Game)

type gameRepoImpl struct {
}

func NewGameRepo() repo.GameRepo {

	return &gameRepoImpl{}
}

func (r *gameRepoImpl) Save(game model.Game) (err error) {
	gameStorage[game.ID()] = game
	return
}

func (r *gameRepoImpl) Find(id uint) (game model.Game, err error) {
	game, ok := gameStorage[id]
	if !ok {
		err = errors.ErrGameNotFound
	}
	return game, err
}
