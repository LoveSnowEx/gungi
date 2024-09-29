package gungi_model

import "image"

type Position interface {
	IsOnBoard() bool
}

type Vector2D struct {
	image.Point
}

type Vector3D struct {
	Vector2D
	Z int
}

type ReservePosition struct {
	Color Color
	Idx   int
}

type DiscardPosition struct {
	Color Color
	Idx   int
}

func (r ReservePosition) IsOnBoard() bool {
	return false
}

func (d DiscardPosition) IsOnBoard() bool {
	return false
}

func NewVector2D(x, y int) Vector2D {
	return Vector2D{Point: image.Point{X: x, Y: y}}
}

func (v Vector2D) Add(w Vector2D) Vector2D {
	return Vector2D{Point: v.Point.Add(w.Point)}
}

func (v Vector2D) Sub(w Vector2D) Vector2D {
	return Vector2D{Point: v.Point.Sub(w.Point)}
}

func (v Vector2D) Eq(w Vector2D) bool {
	return v.Point.Eq(w.Point)
}

func NewVector3D(x, y, z int) Vector3D {
	return Vector3D{Vector2D: NewVector2D(x, y), Z: z}
}

func (v Vector3D) Add(w Vector3D) Vector3D {
	return Vector3D{Vector2D: v.Vector2D.Add(w.Vector2D), Z: v.Z + w.Z}
}

func (v Vector3D) Sub(w Vector3D) Vector3D {
	return Vector3D{Vector2D: v.Vector2D.Sub(w.Vector2D), Z: v.Z - w.Z}
}

func (v Vector3D) Eq(w Vector3D) bool {
	return v.Point.Eq(w.Point) && v.Z == w.Z
}

func (v Vector3D) IsOnBoard() bool {
	return true
}
