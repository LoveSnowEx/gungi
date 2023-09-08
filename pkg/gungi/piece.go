package gungi

import (
	"fmt"
	"log"
)

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
	maxPieceType       int       = iota
)

var (
	_                   Piece = (*basePiece)(nil)
	ErrInvalidPieceType       = fmt.Errorf("invalid piece type")
)

type Piece interface {
	Owner() Player
	Type() PieceType
	Color() Color
	Position() *Position
	CanMoveTo(Position) bool
}

type basePiece struct {
	owner    Player
	type_    PieceType
	color    Color
	position *Position
	board    Board
}

func (t PieceType) Valid() bool {
	val := int(t)
	return val >= 0 && val < maxPieceType
}

func newBasePiece(owner Player, type_ PieceType, board Board) (p *basePiece, err error) {
	if !type_.Valid() {
		err = ErrInvalidPieceType
		return
	}
	p = &basePiece{
		owner: owner,
		type_: type_,
		board: board,
	}
	return
}

func (p basePiece) Type() PieceType {
	return p.type_
}

func (p basePiece) Owner() Player {
	return p.owner
}

func (p basePiece) Color() Color {
	return p.color
}

func (p basePiece) Position() *Position {
	return p.position
}

func (p basePiece) CanMoveTo(pos Position) bool {
	log.Panic("not implemented")
	return false
}
