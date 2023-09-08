package gungi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewPosition(t *testing.T) {
	r := require.New(t)
	p := NewPosition(1, 2, 3)
	r.Equal(1, p.X)
	r.Equal(2, p.Y)
	r.Equal(3, p.Z)
}

func TestNewPosition2D(t *testing.T) {
	r := require.New(t)
	p := NewPosition2D(1, 2)
	r.Equal(1, p.X)
	r.Equal(2, p.Y)
}

func TestPosition2DGet(t *testing.T) {
	r := require.New(t)
	p := NewPosition2D(1, 2)
	x, y := p.Get()
	r.Equal(1, x)
	r.Equal(2, y)
}

func TestPositionGet(t *testing.T) {
	r := require.New(t)
	p := NewPosition(1, 2, 3)
	x, y, z := p.Get()
	r.Equal(1, x)
	r.Equal(2, y)
	r.Equal(3, z)
}
