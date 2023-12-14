package persist

import (
	"context"

	"github.com/LoveSnowEx/gungi/internal/infra/dal"
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
	u := dal.User
	userPo, err := u.WithContext(context.Background()).Where(u.ID.Eq(id)).First()
	if err != nil {
		err = errors.ErrPlayerNotFound
		return
	}
	player = model.NewPlayer(userPo.Name)
	player.SetId(userPo.ID)
	return
}
