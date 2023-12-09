package persist

import (
	"github.com/LoveSnowEx/gungi/internal/infra/po"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/repo"
	"github.com/LoveSnowEx/gungi/pkg/gungi/errors"
)

var (
	_ repo.PlayerRepo = (*playerRepoImpl)(nil)
)

type playerRepoImpl struct {
}

func NewPlayerRepo() repo.PlayerRepo {
	return &playerRepoImpl{}
}

func (r *playerRepoImpl) Find(id uint) (player model.Player, err error) {
	userPo := po.User{}
	if err != nil {
		err = errors.ErrPlayerNotFound
		return
	}
	player = model.NewPlayer(userPo.Name)
	player.SetId(userPo.Id)
	return
}

func (r *playerRepoImpl) Save(player model.Player) (err error) {
	if player == nil {
		err = errors.ErrInvalidPlayer
		return
	}
	userPo := po.User{
		Id:   player.Id(),
		Name: player.Name(),
	}
	return
}
