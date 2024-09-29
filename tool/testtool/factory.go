package testtool

import "github.com/LoveSnowEx/gungi/internal/infra/database"

type Factory[Model any] interface {
	Create(value interface{}) error
}

type factory[Model any] struct {
}

func NewFactory[Model any]() Factory[Model] {
	return &factory[Model]{}
}

func (f *factory[Model]) Create(value interface{}) error {
	return database.Default().Create(value).Error
}
