package gungi

import "image"

type Position2D struct {
	image.Point
}

type Position struct {
	Position2D
	Z int
}

func (p Position2D) Get() (x, y int) {
	return p.X, p.Y
}

func (p Position) Get() (x, y, z int) {
	return p.X, p.Y, p.Z
}
