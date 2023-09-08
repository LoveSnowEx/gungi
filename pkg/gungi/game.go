package gungi

import (
	"fmt"
	"sync"
)

// Color is the color of a player.
type Color int

const (
	WHITE       Color = iota // First player
	BLACK                    // Second player
	PlayerCount int   = iota // Number of players
)

type State int

const (
	STATE_NOT_STARTED State = iota
	STATE_IN_PROGRESS
	STATE_ENDED
)

var (
	_                   Game = (*game)(nil)
	ErrNotEnoughPlayers      = fmt.Errorf("not enough players")
	ErrInvalidColor          = fmt.Errorf("invalid color")
)

// Game represents a game of Gungi.
type Game interface {
	// Board returns the board of the game.
	Board() Board
	// Join adds a player to the game.
	Join(Player, Color) error
	// Players returns the players of the game.
	Players() [PlayerCount]Player
	// CurrentPlayer returns the player whose turn it is.
	CurrentPlayer() Player
	// Start starts the game.
	Start() error
	// State returns the state of the game.
	State() State
	// Actions returns the actions that the current player can take.
	Actions() []Action
	// Apply applies the action to the game.
	Apply(Action) error
}

type game struct {
	players          [PlayerCount]Player
	currentPlayerIdx int
	board            Board
	state            State
	mu               sync.RWMutex
}

func (c Color) Valid() bool {
	val := int(c)
	return val >= 0 && val < PlayerCount
}

func NewGame() Game {
	return &game{
		board: NewBoard(),
	}
}

func (g *game) Board() Board {
	return g.board
}

func (g *game) Join(p Player, c Color) error {
	if !c.Valid() {
		return ErrInvalidColor
	}
	g.players[c] = p
	return nil
}

func (g *game) Players() [PlayerCount]Player {
	return g.players
}

func (g *game) Start() (err error) {
	if g.players[WHITE] == nil || g.players[BLACK] == nil {
		return ErrNotEnoughPlayers
	}
	g.mu.Lock()
	g.state = STATE_IN_PROGRESS
	g.mu.Unlock()
	return
}

func (g *game) State() State {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.state
}

func (g *game) CurrentPlayer() Player {
	return g.players[g.currentPlayerIdx]
}

func (g *game) Actions() []Action {
	return nil
}

func (g *game) Apply(a Action) error {
	return nil
}

func (g *game) endTurn() {
	g.currentPlayerIdx = (g.currentPlayerIdx + 1) % PlayerCount
}
