package gungi_model

type Color interface {
	isColor()
	Other() Color
}

type color uint

const (
	White color = iota
	Black
)

func (c color) isColor() {}

func (c color) Other() Color {
	switch c {
	case White:
		return Black
	case Black:
		return White
	default:
		panic("invalid color")
	}
}

// Colors returns the colors of the game.
func Colors() []Color {
	return []Color{White, Black}
}
