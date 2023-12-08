package persist

import (
	"encoding/json"

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
}

func NewGameRepo() repo.GameRepo {
	return &gameRepoImpl{}
}

func (r *gameRepoImpl) Find(id uint) (game model.Game, err error) {
	jsonBytes, ok := gameStorage[id]
	if !ok {
		err = errors.ErrGameNotFound
		return
	}
	gamePo := po.Game{}
	if err = json.Unmarshal(jsonBytes, &gamePo); err != nil {
		return
	}
	return gamePo.ToModel()
}

func (r *gameRepoImpl) Save(game model.Game) (err error) {
	if game == nil {
		err = errors.ErrInvalidGame
		return
	}
	gamePo := po.NewGame(game)
	if err != nil {
		return
	}
	if gamePo.Id == 0 {
		gamePo.Id = uint(len(gameStorage) + 1)
	}
	jsonBytes, err := json.Marshal(gamePo)
	if err != nil {
		return
	}
	gameStorage[game.Id()] = jsonBytes
	return
}
