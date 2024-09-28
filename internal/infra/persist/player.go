package persist

import (
	"context"

	"github.com/LoveSnowEx/gungi/internal/const/gungi_errors"
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_model"
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_repo"
	"github.com/LoveSnowEx/gungi/internal/infra/dal"
)

var (
	_ gungi_repo.PlayerRepo = (*playerRepoImpl)(nil)
)

type playerRepoImpl struct {
}

func NewPlayerRepo() gungi_repo.PlayerRepo {
	return &playerRepoImpl{}
}

func (r *playerRepoImpl) Find(id uint) (player gungi_model.Player, err error) {
	u := dal.User
	userPo, err := u.WithContext(context.Background()).Where(u.ID.Eq(id)).First()
	if err != nil {
		err = gungi_errors.ErrPlayerNotFound
		return
	}
	player = gungi_model.NewPlayer(userPo.Name)
	player.SetId(userPo.ID)
	return
}
