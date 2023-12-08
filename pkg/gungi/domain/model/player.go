package model

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

type humanPlayer struct {
	basePlayer
}

func NewHumanPlayer(name string) Player {
	return &humanPlayer{
		basePlayer: basePlayer{
			name: name,
		},
	}
}

type aiPlayer struct {
	basePlayer
}

func NewAIPlayer(name string) Player {
	return &aiPlayer{
		basePlayer: basePlayer{
			name: name,
		},
	}
}
