package gungi

import (
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

func newTestPiece(type_ PieceType, owner Player, board Board) Piece {
	p, err := newBasePiece(owner, type_, board)
	if err != nil {
		log.Fatal(err)
	}
	return p
}

func TestBasePiece(t *testing.T) {
	r := require.New(t)

	board := NewBoard()
	owner := newTestHumanPlayer(t)
	for i := range [maxPieceType]struct{}{} {
		type_ := PieceType(i)
		for j := range [PlayerCount]struct{}{} {
			p := newTestPiece(type_, owner, board)
			r.NotNil(p)
			r.Equal(owner, p.Owner())
			r.Equal(type_, p.Type())
		}
	}

	p, err := newBasePiece(owner, type_, board)
	r.NoError(err)
	r.NotNil(p)
	r.Equal(owner, p.Owner())
	r.Equal(type_, p.Type())
	return p
}
