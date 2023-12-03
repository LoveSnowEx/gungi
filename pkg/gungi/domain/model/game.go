package model

import "github.com/LoveSnowEx/gungi/pkg/gungi/infra/errors"

type Phase int

const (
	Setup Phase = iota
	Play
)

var (
	PieceAmounts = map[PieceType]int{
		Marshal:           1,
		General:           1,
		LieutenantGeneral: 1,
		MajorGeneral:      2,
		Samurai:           2,
		Lancer:            3,
		Knight:            2,
		Spy:               2,
		Fortress:          2,
		Pawn:              4,
		Cannon:            1,
		Musketeer:         1,
		Archer:            2,
		Captain:           1,
	}
)

type Game interface {
	// ID returns the ID of the game.
	ID() uint
	// SetID sets the ID of the game.
	SetID(id uint)
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

func (g game) ID() uint {
	return g.id
}

func (g *game) SetID(id uint) {
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
		if g.Player(c).ID() == player.ID() {
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
		if g.Player(c).ID() == player.ID() {
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
