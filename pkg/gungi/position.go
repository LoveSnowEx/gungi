package gungi

import "image"

type Position2D struct {
	image.Point
}

type Position struct {
	Position2D
	Z int
}

func NewPosition2D(x, y int) Position2D {
	return Position2D{image.Point{x, y}}
}

func NewPosition(x, y, z int) Position {
	return Position{NewPosition2D(x, y), z}
}

func (p Position2D) Get() (x, y int) {
	return p.X, p.Y
}

func (p Position) Get() (x, y, z int) {
	return p.X, p.Y, p.Z
}
