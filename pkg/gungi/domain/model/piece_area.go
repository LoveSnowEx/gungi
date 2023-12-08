package model

import (
	"github.com/LoveSnowEx/gungi/pkg/gungi/errors"
	"k8s.io/apimachinery/pkg/util/sets"
)

type PieceArea interface {
	// Add adds a piece to the area.
	Add(piece Piece) error
	// Remove removes a piece from the area.
	Remove(piece Piece) error
	// Contains returns whether the area contains the piece.
	Contains(piece Piece) bool
	// Pieces returns the pieces in the area.
	Pieces() []Piece
}

type pieceArea struct {
	pieces sets.Set[Piece]
}

func NewPieceArea() PieceArea {
	return &pieceArea{
		sets.New[Piece](),
	}
}

func (a *pieceArea) Add(piece Piece) (err error) {
	if a.Contains(piece) {
		return errors.ErrPieceAlreadyExists
	}
	a.pieces.Insert(piece)
	return
}

func (a *pieceArea) Remove(piece Piece) (err error) {
	if !a.Contains(piece) {
		return errors.ErrPieceNotFound
	}
	a.pieces.Delete(piece)
	return
}

func (a *pieceArea) Contains(piece Piece) bool {
	return a.pieces.Has(piece)
}

func (a *pieceArea) Pieces() []Piece {
	return a.pieces.UnsortedList()
}
