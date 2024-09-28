package testtool

type Factory[Po any] interface {
	Create(values ...*Po) error
}

type Do[Po any] interface {
	Create(values ...*Po) error
}

type factory[Po any] struct {
	newDo func() Do[Po]
}

func New[Po any](newDo func() Do[Po]) Factory[Po] {
	return &factory[Po]{newDo: newDo}
}

func (f *factory[Po]) Create(values ...*Po) error {
	return f.newDo().Create(values...)
}
