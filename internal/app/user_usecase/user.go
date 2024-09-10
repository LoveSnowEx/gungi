package user_usecase

import (
	"github.com/LoveSnowEx/gungi/internal/domain/user_model"
	"github.com/LoveSnowEx/gungi/internal/domain/user_repo"
)

type Config struct {
	UserRepo user_repo.Repo
}

type Usecase interface {
	Find(id uint) (user user_model.User, err error)
	Create(name string) (user user_model.User, err error)
}

type usecase struct {
	userRepo user_repo.Repo
}

func New(config *Config) Usecase {
	return &usecase{
		userRepo: config.UserRepo,
	}
}

func (u *usecase) Find(id uint) (user user_model.User, err error) {
	return u.userRepo.Find(id)
}

func (u *usecase) Create(name string) (user user_model.User, err error) {
	user = user_model.NewUser(name)
	err = u.userRepo.Create(user)
	if err != nil {
		user = nil
		return
	}
	return

}
