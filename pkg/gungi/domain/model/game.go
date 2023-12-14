package model

import "github.com/LoveSnowEx/gungi/pkg/gungi/errors"

type Phase interface {
	isPhase()
}

type phase uint

func (p phase) isPhase() {}

const (
	Setup phase = iota
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
	Board() Board
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
	// CurrentTurn returns the color of the player whose turn it is.
	CurrentTurn() Color
	// SetCurrentTurn sets the color of the player whose turn it is.
	SetCurrentTurn(color Color)
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

func NewGame() (g Game) {
	g = &game{
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
	return
}

func (g game) Id() uint {
	return g.id
}

func (g *game) SetId(id uint) {
	g.id = id
}

func (g game) Board() Board {
	return g.board
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
			err = errors.ErrPlayerAlreadyJoined
			return
		}
	}
	if g.Player(color) != nil {
		err = errors.ErrTeamFull
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
	err = errors.ErrPlayerNotFound
	return
}

func (g *game) SetCurrentTurn(color Color) {
	g.currentTurn = color
}

func (g game) CurrentTurn() Color {
	return g.currentTurn
}

func (g *game) SetPhase(phase Phase) {
	g.phase = phase
}

func (g game) Phase() Phase {
	return g.phase
}
