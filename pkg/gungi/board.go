package gungi

const maxStackHeight = 3

type Board interface {
	// Get returns the piece at the given position.
	Get(Position) Piece
	// Highest returns the height of the stack at the given position.
	Highest(Position2D) int
}

type stack []Piece

// The board is a 9x9 grid of stacks of pieces.
// The first index is the x coordinate, the second is the y coordinate.
// The top of the stack is the last element of the slice.
type board [9][9]stack

func NewBoard() Board {
	b := board{}
	for x := range b {
		for y := range b[x] {
			b[x][y] = make(stack, 0, maxStackHeight)
		}
	}
	return b
}

func (b board) Get(p Position) Piece {
	x, y, z := p.Get()
	stack := b.getStack(p.Position2D)
	if z >= len(stack) {
		return nil
	}
	return b[x][y][z]
}

func (b board) Highest(p Position2D) int {
	return len(b.getStack(p))
}

func (b board) getStack(p Position2D) stack {
	x, y := p.Get()
	return b[x][y]
}
