package gungi_model

type Player interface {
	Id() uint
	SetId(id uint)
	Name() string
	SetName(name string)
}

type basePlayer struct {
	id   uint
	name string
}

func (p basePlayer) Id() uint {
	return p.id
}

func (p *basePlayer) SetId(id uint) {
	p.id = id
}

func (p basePlayer) Name() string {
	return p.name
}

func (p *basePlayer) SetName(name string) {
	p.name = name
}

type player struct {
	basePlayer
}

func NewPlayer(userId uint, name string) Player {
	return &player{
		basePlayer: basePlayer{
			name: name,
		},
	}
}
