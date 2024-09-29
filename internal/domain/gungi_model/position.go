package gungi_model

type BoardPosition = Vector3D

type ReservePosition struct {
	Color Color
	Idx   uint
}

type DiscardPosition struct {
	Color Color
	Idx   uint
}

func NewBoardPosition(x, y, z int) BoardPosition {
	return BoardPosition{
		Vector2D: NewVector2D(int(x), int(y)),
		Z:        int(z),
	}
}

func NewReservePosition(color Color, idx uint) ReservePosition {
	return ReservePosition{
		Color: color,
		Idx:   idx,
	}
}

func NewDiscardPosition(color Color, idx uint) DiscardPosition {
	return DiscardPosition{
		Color: color,
		Idx:   idx,
	}
}
