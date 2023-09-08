package gungi

const playerCount = 2

// Game represents a game of Gungi.
type Game interface {
	// Board returns the board of the game.
	Board() Board
	// Players returns the players of the game.
	Players() [playerCount]Player
	// CurrentPlayer returns the player whose turn it is.
	CurrentPlayer() Player
	// Start starts the game.
	Start()
	// Actions returns the actions that the current player can take.
	Actions() []Action
	// Apply applies the action to the game.
	Apply(Action) error
}

type game struct {
	players          [playerCount]Player
	currentPlayerIdx int
	board            Board
}

func NewGame() Game {
	return &game{}
}

func (g *game) Board() Board {
	return g.board
}

func (g *game) Players() [playerCount]Player {
	return g.players
}

func (g *game) Start() {
	g.board = NewBoard()
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
	g.currentPlayerIdx = (g.currentPlayerIdx + 1) % len(g.players)
}
