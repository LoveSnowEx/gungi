package gungi

// Player represents a player in the game.
type Player interface {
	// Pieces returns the pieces that the player has on the board.
	Pieces() []Piece
	// Hand returns the pieces that the player has in their hand.
	Hand() []Piece
}

type player struct {
	pieces []Piece
	hand   []Piece
}

type humanPlayer struct {
	player
}

func (p *player) Pieces() []Piece {
	return p.pieces
}

func (p *player) Hand() []Piece {
	return p.hand
}

func NewHumanPlayer() (Player, error) {
	return &humanPlayer{}, nil
}
