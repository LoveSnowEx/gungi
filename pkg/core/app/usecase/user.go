package usecase

import (
	"github.com/LoveSnowEx/gungi/pkg/core/domain/model"
	"github.com/LoveSnowEx/gungi/pkg/core/domain/repo"
)

type UserUsecaseConfig struct {
	UserRepo repo.UserRepo
}

type UserUsecase interface {
	Find(id uint) (user model.User, err error)
	Create(name string) (user model.User, err error)
}

type userUsecase struct {
	userRepo repo.UserRepo
}

func NewUserUsecase(config *UserUsecaseConfig) UserUsecase {
	return &userUsecase{
		userRepo: config.UserRepo,
	}
}

func (u *userUsecase) Find(id uint) (user model.User, err error) {
	return u.userRepo.Find(id)
}

func (u *userUsecase) Create(name string) (user model.User, err error) {
	user = model.NewUser(name)
	err = u.userRepo.Create(user)
	if err != nil {
		user = nil
		return
	}
	return

}
