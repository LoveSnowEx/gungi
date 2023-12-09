package persist

import (
	"context"

	"github.com/LoveSnowEx/gungi/internal/infra/dal"
	"github.com/LoveSnowEx/gungi/internal/infra/po"
	"github.com/LoveSnowEx/gungi/pkg/core/domain/model"
	"github.com/LoveSnowEx/gungi/pkg/core/errors"
	"github.com/LoveSnowEx/gungi/pkg/core/repo"
)

var (
	_ repo.UserRepo = (*userRepoImpl)(nil)
)

type userRepoImpl struct {
}

func NewUserRepo() repo.UserRepo {
	return &userRepoImpl{}
}

func (r *userRepoImpl) Find(id uint) (user model.User, err error) {
	u := dal.User
	userPo, err := u.WithContext(context.Background()).Where(u.ID.Eq(id)).First()
	if err != nil {
		err = errors.ErrUserNotFound
	}
	user = model.NewUser(userPo.Name)
	user.SetId(uint(userPo.ID))
	return
}

func (r *userRepoImpl) Create(name string) (user model.User, err error) {
	userPo := po.User{
		Name: name,
	}
	u := dal.User
	err = u.WithContext(context.Background()).Create(&userPo)
	return
}
