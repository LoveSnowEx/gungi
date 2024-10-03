package po

import (
	"github.com/uptrace/bun"
)

type Game struct {
	bun.BaseModel `bun:"table:games,alias:g"`
	TimeStamp
	ID      uint     `bun:",pk,autoincrement"`
	Players []Player `bun:",rel:has-many"`
	Pieces  []Piece  `bun:",rel:has-many"`
	Turn    uint     `bun:",nullzero,notnull,default:0"`
	Phase   uint     `bun:",nullzero,notnull,default:0"`
	Winner  *uint    `bun:""`
}

type Piece struct {
	bun.BaseModel `bun:"table:pieces"`
	GameID        uint `bun:",pk"`
	Type          uint `bun:",notnull"`
	Color         uint `bun:",notnull"`
	Position      uint `bun:",notnull"`
}

type Player struct {
	bun.BaseModel `bun:"table:players"`
	GameID        uint `bun:",pk"`
	UserID        uint `bun:",pk"`
	User          User `bun:",rel:belongs-to"`
	Color         uint `bun:",notnull"`
}
