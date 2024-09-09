package user_usecase

import (
	"github.com/LoveSnowEx/gungi/internal/domain/user_model"
	"github.com/LoveSnowEx/gungi/internal/domain/user_repo"
)

type UserUsecaseConfig struct {
	UserRepo user_repo.UserRepo
}

type UserUsecase interface {
	Find(id uint) (user user_model.User, err error)
	Create(name string) (user user_model.User, err error)
}

type userUsecase struct {
	userRepo user_repo.UserRepo
}

func NewUserUsecase(config *UserUsecaseConfig) UserUsecase {
	return &userUsecase{
		userRepo: config.UserRepo,
	}
}

func (u *userUsecase) Find(id uint) (user user_model.User, err error) {
	return u.userRepo.Find(id)
}

func (u *userUsecase) Create(name string) (user user_model.User, err error) {
	user = user_model.NewUser(name)
	err = u.userRepo.Create(user)
	if err != nil {
		user = nil
		return
	}
	return

}
