package gungi_model

const (
	BoardRows   = 9
	BoardCols   = 9
	BoardLevels = 3
)

type Board interface {
	Board() [9][9][3]Piece
	Get(loc Vector3D) Piece
	Set(loc Vector3D, piece Piece)
	Take(loc Vector3D) Piece
}

type board struct {
	board [BoardRows][BoardCols][BoardLevels]Piece
}

func NewBoard() *board {
	return &board{
		board: [BoardRows][BoardCols][BoardLevels]Piece{},
	}
}

func (b board) Board() [9][9][3]Piece {
	return b.board
}

func (b board) Get(loc Vector3D) Piece {
	return b.board[loc.X][loc.Y][loc.Z]
}

func (b *board) Set(loc Vector3D, piece Piece) {
	b.board[loc.X][loc.Y][loc.Z] = piece
}

func (b *board) Take(loc Vector3D) (piece Piece) {
	piece = b.Get(loc)
	b.board[loc.X][loc.Y][loc.Z] = nil
	return
}
