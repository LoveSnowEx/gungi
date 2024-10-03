package gungi_model

type Color uint

const (
	Black Color = iota
	White
)

func (c Color) Opposite() Color {
	return c ^ 1
}

// Colors returns a slice of all possible colors.
func Colors() []Color {
	return []Color{Black, White}
}
