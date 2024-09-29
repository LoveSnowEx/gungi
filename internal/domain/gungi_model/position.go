package gungi_model

type ReservePosition struct {
	Color Color
	Idx   uint
}

type DiscardPosition struct {
	Color Color
	Idx   uint
}

type BoardPosition = Vector3D
