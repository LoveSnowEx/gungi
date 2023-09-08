package gungi

type Action interface {
	// Apply applies the action to the game.
	Apply(*Game) error
}

type MoveAction struct {
	Piece    Piece
	Position Position
}

func (a MoveAction) Apply(g *Game) error {
	return nil
}

type PlaceAction struct {
	Piece    Piece
	Position Position2D
}

func (a PlaceAction) Apply(g *Game) error {
	return nil
}
