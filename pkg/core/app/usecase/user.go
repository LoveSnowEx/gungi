package usecase

import (
	"github.com/LoveSnowEx/gungi/pkg/core/domain/model"
	"github.com/LoveSnowEx/gungi/pkg/core/repo"
)

type UserUseCase interface {
	Find(id uint) (user model.User, err error)
	Create(name string) (user model.User, err error)
}

type userUseCase struct {
	userRepo repo.UserRepo
}

func NewUserUseCase(userRepo repo.UserRepo) UserUseCase {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u *userUseCase) Find(id uint) (user model.User, err error) {
	return u.userRepo.Find(id)
}

func (u *userUseCase) Create(name string) (user model.User, err error) {
	return u.userRepo.Create(name)
}
