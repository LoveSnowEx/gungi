package gungi

import (
	"testing"

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
	r := require.New(t)
	game := newTestGame(t)
	r.NotNil(game.Board())
}

func TestNoGamePlayers(t *testing.T) {
	r := require.New(t)
	game := newTestGame(t)
	r.Empty(game.Players())
	r.Nil(game.CurrentPlayer())
}

func TestGameJoin(t *testing.T) {
	r := require.New(t)
	game := newTestGame(t)
	p1 := newTestHumanPlayer(t)
	p2 := newTestHumanPlayer(t)
	r.NoError(game.Join(p1, WHITE))
	r.NoError(game.Join(p2, BLACK))
	r.Equal([playerCount]Player{p1, p2}, game.Players())
	r.Equal(p1, game.CurrentPlayer())
}
