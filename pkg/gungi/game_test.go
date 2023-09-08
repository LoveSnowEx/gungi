package gungi

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestGame(t *testing.T) Game {
	r := require.New(t)
	game, err := NewGame()
	r.NoError(err)
	r.NotNil(game)
	return game
}

func TestNewGame(t *testing.T) {
	newTestGame(t)
}

func TestGameBoard(t *testing.T) {
	a := assert.New(t)
	game := newTestGame(t)
	a.NotNil(game.Board())
}

func TestNoGamePlayers(t *testing.T) {
	a := assert.New(t)
	game := newTestGame(t)
	a.Empty(game.Players())
	a.Nil(game.CurrentPlayer())
}

func TestGameJoin(t *testing.T) {
	a := assert.New(t)
	game := newTestGame(t)
	p1 := newTestHumanPlayer(t)
	p2 := newTestHumanPlayer(t)
	a.NoError(game.Join(p1, WHITE))
	a.NoError(game.Join(p2, BLACK))
	a.Equal([playerCount]Player{p1, p2}, game.Players())
	a.Equal(p1, game.CurrentPlayer())
}
