package gungi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func newTestHumanPlayer(t *testing.T) Player {
	r := require.New(t)
	player, err := NewHumanPlayer()
	r.NoError(err)
	r.NotNil(player)
	return player
}

func TestNewHumanPlayer(t *testing.T) {
	newTestHumanPlayer(t)
}
