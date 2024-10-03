package gungi_model

const AreaSize = 25

type PieceArea interface {
	Area() [AreaSize]Piece
	Get(idx uint) Piece
	Set(idx uint, piece Piece)
}

type pieceArea struct {
	pieces [AreaSize]Piece
}

func NewPieceArea() *pieceArea {
	return &pieceArea{
		pieces: [AreaSize]Piece{},
	}
}

func (pa pieceArea) Area() [AreaSize]Piece {
	return pa.pieces
}

func (pa pieceArea) Get(idx uint) Piece {
	return pa.pieces[idx]
}

func (pa *pieceArea) Set(idx uint, piece Piece) {
	pa.pieces[idx] = piece
}
