package user_model

type User interface {
	Id() uint
	SetId(id uint)
	Name() string
	SetName(name string)
}

type user struct {
	id   uint
	name string
}

func NewUser(name string) User {
	return &user{
		name: name,
	}
}

func (u *user) Id() uint {
	return u.id
}

func (u *user) SetId(id uint) {
	u.id = id
}

func (u *user) Name() string {
	return u.name
}

func (u *user) SetName(name string) {
	u.name = name
}
