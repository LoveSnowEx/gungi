package model

type Vector2D interface {
	X() int
	Y() int
}

type Vector3D interface {
	Vector2D
	Z() int
}

type Vector2DImpl struct {
	x int
	y int
}

type Vector3DImpl struct {
	Vector2DImpl
	z int
}

func NewVector2D(x, y int) Vector2D {
	return Vector2DImpl{x, y}
}

func (v Vector2DImpl) X() int {
	return v.x
}

func (v Vector2DImpl) Y() int {
	return v.y
}

func NewVector3D(x, y, z int) Vector3D {
	return Vector3DImpl{Vector2DImpl{x, y}, z}
}

func (v Vector3DImpl) Z() int {
	return v.z
}
