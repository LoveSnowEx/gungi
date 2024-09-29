package gungi_model

import "github.com/LoveSnowEx/gungi/internal/const/gungi_errors"

type Phase uint

const (
	Setup Phase = iota
	Prepare
	Play
	End
)

type Game interface {
	// Id returns the id of the game.
	Id() uint
	// SetId sets the id of the game.
	SetId(id uint)
	// Board returns the board of the game.
	Board() [BoardRows][BoardCols][BoardLevels]Piece
	// Reserve returns the reserve of the game.
	Reserve(color Color) PieceArea
	// Discard returns the discard area of the game.
	Discard(color Color) PieceArea
	// Player returns the player of the game.
	Player(color Color) Player
	// Join adds a player to the game.
	Join(color Color, player Player) error
	// Leave removes a player from the game.
	Leave(player Player) error
	// Turn returns the color of the player whose turn it is.
	Turn() Color
	// SetTurn sets the color of the player whose turn it is.
	SetTurn(color Color)
	// Phase returns the current phase of the game.
	Phase() Phase
	// SetPhase sets the current phase of the game.
	SetPhase(phase Phase)
}

type game struct {
	id          uint
	board       Board
	reserve     map[Color]PieceArea
	discard     map[Color]PieceArea
	players     map[Color]Player
	currentTurn Color
	phase       Phase
}

func NewGame() *game {
	return &game{
		board: NewBoard(),
		reserve: map[Color]PieceArea{
			White: NewPieceArea(),
			Black: NewPieceArea(),
		},
		discard: map[Color]PieceArea{
			White: NewPieceArea(),
			Black: NewPieceArea(),
		},
		players: map[Color]Player{
			White: nil,
			Black: nil,
		},
		currentTurn: White,
		phase:       Setup,
	}
}

func (g game) Id() uint {
	return g.id
}

func (g *game) SetId(id uint) {
	g.id = id
}

func (g game) Board() [BoardRows][BoardCols][BoardLevels]Piece {
	return g.board.Board()
}

func (g game) Reserve(color Color) PieceArea {
	return g.reserve[color]
}

func (g game) Discard(color Color) PieceArea {
	return g.discard[color]
}

func (g game) Player(color Color) Player {
	player, ok := g.players[color]
	if !ok {
		return nil
	}
	return player
}

func (g *game) Join(color Color, player Player) (err error) {
	for _, c := range Colors() {
		if g.Player(c) != nil && g.Player(c).Id() == player.Id() {
			err = gungi_errors.ErrPlayerAlreadyJoined
			return
		}
	}
	if g.Player(color) != nil {
		err = gungi_errors.ErrTeamFull
		return
	}
	g.players[color] = player
	return
}

func (g *game) Leave(player Player) (err error) {
	for _, c := range Colors() {
		if g.Player(c).Id() == player.Id() {
			g.players[c] = nil
			return
		}
	}
	err = gungi_errors.ErrPlayerNotFound
	return
}

func (g game) Turn() Color {
	return g.currentTurn
}

func (g *game) SetTurn(color Color) {
	g.currentTurn = color
}

func (g game) Phase() Phase {
	return g.phase
}

func (g *game) SetPhase(phase Phase) {
	g.phase = phase
}
