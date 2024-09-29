package persist

import (
	"github.com/LoveSnowEx/gungi/internal/const/user_errors"
	"github.com/LoveSnowEx/gungi/internal/domain/user_model"
	"github.com/LoveSnowEx/gungi/internal/domain/user_repo"

	"github.com/LoveSnowEx/gungi/internal/infra/database"
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
	userPo := po.User{}
	if err = database.Default().Where("id = ?", id).First(&userPo).Error; err != nil {
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
	if err = database.Default().Create(&userPo).Error; err != nil {
		return
	}
	id = userPo.ID
	return
}
