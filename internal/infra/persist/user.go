package persist

import (
	"context"

	"github.com/LoveSnowEx/gungi/internal/const/user_errors"
	"github.com/LoveSnowEx/gungi/internal/domain/user_model"
	"github.com/LoveSnowEx/gungi/internal/domain/user_repo"
	"github.com/LoveSnowEx/gungi/internal/infra/dal"
	"github.com/LoveSnowEx/gungi/internal/infra/po"
)

var (
	_ user_repo.Repo = (*userRepoImpl)(nil)
)

type userRepoImpl struct {
}

func NewUserRepo() user_repo.Repo {
	return &userRepoImpl{}
}

func (r *userRepoImpl) Find(id uint) (user user_model.User, err error) {
	u := dal.User
	userPo, err := u.WithContext(context.Background()).Where(u.ID.Eq(id)).First()
	if err != nil {
		err = user_errors.ErrUserNotFound
		return
	}
	user = user_model.NewUser(userPo.Name)
	user.SetId(userPo.ID)
	return
}

func (r *userRepoImpl) Create(user user_model.User) (err error) {
	userPo := po.User{
		Name: user.Name(),
	}
	u := dal.User
	err = u.WithContext(context.Background()).Create(&userPo)
	if err != nil {
		return
	}
	user.SetId(userPo.ID)
	return
}
