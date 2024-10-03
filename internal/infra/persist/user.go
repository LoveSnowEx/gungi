package persist

import (
	"context"

	"github.com/LoveSnowEx/gungi/internal/const/user_errors"
	"github.com/LoveSnowEx/gungi/internal/domain/user_model"
	"github.com/LoveSnowEx/gungi/internal/domain/user_repo"
	"github.com/uptrace/bun"

	"github.com/LoveSnowEx/gungi/internal/infra/po"
)

var (
	_ user_repo.Repo = (*userRepo)(nil)
)

type userRepo struct {
	db *bun.DB
}

func NewUserRepo(db *bun.DB) *userRepo {
	_, err := db.NewCreateTable().
		IfNotExists().
		Model(&po.User{}).
		Exec(context.Background())
	if err != nil {
		panic(err)
	}
	return &userRepo{db}
}

func (r *userRepo) Find(id uint) (user *user_model.User, err error) {
	userPo := po.User{}
	err = r.db.NewSelect().
		Model(&po.User{}).
		Where("id = ?", id).
		Scan(context.Background())
	if err != nil {
		err = user_errors.ErrUserNotFound
		return
	}
	user = user_model.NewUser(userPo.Name)
	user.ID = userPo.ID
	return
}

func (r *userRepo) Save(user *user_model.User) (err error) {
	if user.ID == 0 {
		return r.create(user)
	}
	return r.update(user)
}

func (r *userRepo) create(user *user_model.User) (err error) {
	userPo := po.User{
		Name: user.Name,
	}
	_, err = r.db.NewInsert().
		Model(&userPo).
		Exec(context.Background())
	if err != nil {
		return
	}
	user.ID = userPo.ID
	return
}

func (r *userRepo) update(user *user_model.User) (err error) {
	userPo := po.User{
		ID:   user.ID,
		Name: user.Name,
	}
	_, err = r.db.NewUpdate().
		Model(&userPo).
		WherePK().
		Set("name = ?", user.Name).
		Exec(context.Background())
	return
}
