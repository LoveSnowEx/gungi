package gungi_model

const (
	BoardRows   = 9
	BoardCols   = 9
	BoardLevels = 3
)

type Board interface {
	Board() [9][9][3]Piece
	Get(loc Vector3D) Piece
	GetLocation(piece Piece) Vector3D
	Set(loc Vector3D, piece Piece)
	Take(piece Piece)
}

type board struct {
	board     [BoardRows][BoardCols][BoardLevels]Piece
	locations map[Piece]Vector3D
}

func NewBoard() *board {
	return &board{
		board:     [BoardRows][BoardCols][BoardLevels]Piece{},
		locations: map[Piece]Vector3D{},
	}
}

func (b board) Board() [9][9][3]Piece {
	return b.board
}

func (b board) Get(loc Vector3D) Piece {
	return b.board[loc.X][loc.Y][loc.Z]
}

func (b board) GetLocation(piece Piece) Vector3D {
	return b.locations[piece]
}

func (b *board) Set(loc Vector3D, piece Piece) {
	b.board[loc.X][loc.Y][loc.Z] = piece
	b.locations[piece] = loc
}

func (b *board) Take(piece Piece) {
	loc := b.locations[piece]
	b.board[loc.X][loc.Y][loc.Z] = nil
	delete(b.locations, piece)
}
