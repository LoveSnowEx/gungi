package user_repo

import "github.com/LoveSnowEx/gungi/internal/domain/user_model"

type Repo interface {
	Find(id uint) (user *user_model.User, err error)
	Save(user *user_model.User) (err error)
}
