package persist

import (
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/repo"
	"github.com/LoveSnowEx/gungi/pkg/gungi/errors"
	"github.com/LoveSnowEx/gungi/pkg/gungi/infra/po"
)

var (
	_           repo.GameRepo = (*gameRepoImpl)(nil)
	gameStorage               = make(map[uint][]byte)
)

type gameRepoImpl struct {
	gameMarshaller po.GameMarshaller
}

func NewGameRepo() repo.GameRepo {
	return &gameRepoImpl{
		gameMarshaller: po.NewGameMarshaller(),
	}
}

func (r *gameRepoImpl) Find(id uint) (game model.Game, err error) {
	jsonBytes, ok := gameStorage[id]
	if !ok {
		err = errors.ErrGameNotFound
		return
	}
	return r.gameMarshaller.Unmarshal(jsonBytes)
}

func (r *gameRepoImpl) Save(game model.Game) (err error) {
	if game == nil {
		err = errors.ErrInvalidGame
		return
	}
	jsonBytes, err := r.gameMarshaller.Marshal(game)
	gameStorage[game.Id()] = jsonBytes
	return
}
