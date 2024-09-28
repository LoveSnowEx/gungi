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
	_ user_repo.Repo = (*userRepo)(nil)
)

type userRepo struct {
}

func NewUserRepo() *userRepo {
	return &userRepo{}
}

func (r *userRepo) Find(id uint) (user user_model.User, err error) {
	u := dal.User
	userPo, err := u.WithContext(context.Background()).Where(u.ID.Eq(id)).First()
	if err != nil {
		err = user_errors.ErrUserNotFound
		return
	}
	user = user_model.NewUser(userPo.Name)
	user.ID = userPo.ID
	return
}

func (r *userRepo) Create(user user_model.User) (id uint, err error) {
	userPo := po.User{
		Name: user.Name,
	}
	u := dal.User
	err = u.WithContext(context.Background()).Create(&userPo)
	if err != nil {
		return
	}
	id = userPo.ID
	return
}
