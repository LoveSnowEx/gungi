package gungi

import "log"

type PieceType int

const (
	MARSHAL            PieceType = iota // 帥
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
	Type() PieceType
	Position() *Position
	CanMoveTo(Position) bool
}

type basePiece struct {
	owner    Player
	type_    PieceType
	position *Position
	board    board
}

func newBasePiece(owner Player, type_ PieceType, position *Position, board board) basePiece {
	return basePiece{
		owner:    owner,
		type_:    type_,
		position: position,
		board:    board,
	}
}

func (p basePiece) Type() PieceType {
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
