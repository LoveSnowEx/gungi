package gungi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func newTestBoard(t *testing.T) Board {
	r := require.New(t)
	board := NewBoard()
	r.NotNil(board)
	return board
}

func TestNewBoard(t *testing.T) {
	newTestBoard(t)
}

func TestEmptyBoardGet(t *testing.T) {
	r := require.New(t)
	board := newTestBoard(t)
	z := 0
	for x := range [9]struct{}{} {
		for y := range [9]struct{}{} {
			pos := NewPosition(x, y, z)
			r.Nil(board.Get(pos))
		}
	}
}

func TestEmptyBoardHighest(t *testing.T) {
	r := require.New(t)
	board := newTestBoard(t)
	for x := range [9]struct{}{} {
		for y := range [9]struct{}{} {
			pos := NewPosition2D(x, y)
			r.Equal(0, board.Highest(pos))
		}
	}
}
