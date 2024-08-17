package repo

import "github.com/LoveSnowEx/gungi/pkg/core/domain/model"

type UserRepo interface {
	Find(id uint) (user model.User, err error)
	Create(user model.User) (err error)
}
