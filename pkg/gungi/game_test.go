package gungi

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func newTestGame(t *testing.T) Game {
	r := require.New(t)
	game := NewGame()
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

func TestGamePlayers(t *testing.T) {
	r := require.New(t)
	game := newTestGame(t)
	p1 := newTestHumanPlayer(t)
	p2 := newTestHumanPlayer(t)
	_ = game.Join(p1, WHITE)
	_ = game.Join(p2, BLACK)
	r.Equal([PlayerCount]Player{p1, p2}, game.Players())
	r.Equal(p1, game.CurrentPlayer())
}

func TestGameJoin(t *testing.T) {
	r := require.New(t)
	game := newTestGame(t)
	p1 := newTestHumanPlayer(t)
	p2 := newTestHumanPlayer(t)
	r.NoError(game.Join(p1, WHITE))
	r.NoError(game.Join(p2, BLACK))
	r.Equal([PlayerCount]Player{p1, p2}, game.Players())
	r.Equal(p1, game.CurrentPlayer())

	// Test that you can't join a game with wrong color.
	p3 := newTestHumanPlayer(t)
	r.Error(game.Join(p3, -1))
	r.Error(game.Join(p3, PlayerCount))

	// Test that you can't join a game that has already started.
	stratedCh := make(chan struct{})
	// Start the game in a goroutine so we can wait for it to start.
	go func() {
		_ = game.Start()
	}()
	// Wait for the game to start.
	go func() {
		for game.State() != STATE_IN_PROGRESS {
			time.Sleep(10 * time.Millisecond)
		}
		stratedCh <- struct{}{}
	}()
	<-stratedCh

	r.Error(game.Join(p3, WHITE))
	r.Error(game.Join(p3, BLACK))
}
