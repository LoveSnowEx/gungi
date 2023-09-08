package gungi

import "log"

type PiceeType int

const (
	MARSHAL            PiceeType = iota // 帥
	GENERAL                             // 大
	LIEUTENANT_GENERAL                  // 中
	MAJOR_GENERAL                       // 小
	SAMURAI                             // 侍
	LANCER                              // 槍
	KNIGHT                              // 馬
	SPY                                 // 忍
	FORTRESS                            // 砦
	PAWN                                // 兵
	CANNON                              // 砲
	MUSKETEER                           // 筒
	ARCHER                              // 弓
	CAPTAIN                             // 謀
)

type Piece interface {
	Owner() Player
	Type() PiceeType
	Position() *Position
	CanMoveTo(Position) bool
}

type basePiece struct {
	owner    Player
	type_    PiceeType
	position *Position
	board    board
}

func newBasePiece(owner Player, type_ PiceeType, position *Position, board board) basePiece {
	return basePiece{
		owner:    owner,
		type_:    type_,
		position: position,
		board:    board,
	}
}

func (p basePiece) Type() PiceeType {
	return p.type_
}

func (p basePiece) Owner() Player {
	return p.owner
}

func (p basePiece) Position() *Position {
	return p.position
}

func (p basePiece) CanMoveTo(pos Position) bool {
	log.Panic("not implemented")
	return false
}
