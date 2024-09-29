package persist

import (
	"github.com/LoveSnowEx/gungi/internal/const/gungi_errors"
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_model"
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_repo"
	"github.com/LoveSnowEx/gungi/internal/infra/database"
	"github.com/LoveSnowEx/gungi/internal/infra/po"
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
	playerPo := po.User{}
	err = database.Default().Where("id = ?", id).First(&playerPo).Error
	if err != nil {
		err = gungi_errors.ErrPlayerNotFound
		return
	}
	player = gungi_model.NewPlayer(playerPo.Name)
	player.SetId(playerPo.ID)
	return
}
