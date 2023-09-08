package gungi

const PlayerCount = 2

var (
	_ Game = (*game)(nil)
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
	Start()
	// Actions returns the actions that the current player can take.
	Actions() []Action
	// Apply applies the action to the game.
	Apply(Action) error
}

type game struct {
	players          [PlayerCount]Player
	currentPlayerIdx int
	board            Board
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
	g.players[c] = p
	return nil
}

func (g *game) Players() [PlayerCount]Player {
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
