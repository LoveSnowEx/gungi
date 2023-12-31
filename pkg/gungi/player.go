package gungi

import "fmt"

var (
	_ Player = (*HumanPlayer)(nil)

	ErrNilPlayer = fmt.Errorf("nil player")
)

// Player represents a player in the game.
type Player interface {
	// Pieces returns the pieces that the player has on the board.
	Pieces() []Piece
	// Hand returns the pieces that the player has in their hand.
	Hand() []Piece
	// DoTurn performs the player's turn.
	DoTurn()
}

type player struct {
	pieces []Piece
	hand   []Piece
}

type HumanPlayer struct {
	player
}

func (p *player) Pieces() []Piece {
	return p.pieces
}

func (p *player) Hand() []Piece {
	return p.hand
}

func NewHumanPlayer() (p *HumanPlayer, err error) {
	// TODO
	p = &HumanPlayer{}
	return
}

func (p *HumanPlayer) DoTurn() {
	// TODO
}
