package user_repo

import "github.com/LoveSnowEx/gungi/internal/domain/user_model"

type UserRepo interface {
	Find(id uint) (user user_model.User, err error)
	Create(user user_model.User) (err error)
}
