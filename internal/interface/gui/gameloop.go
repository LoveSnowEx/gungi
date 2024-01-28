package gui

import (
	"fmt"
	"image/color"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

func (g *game) Update() error {
	cnt++
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	lines := make([]string, 10)
	lines[0] = fmt.Sprintf("FPS: %f, TPS: %d", ebiten.ActualFPS(), ebiten.TPS())
	lines[1] = fmt.Sprintf("cnt: %d", cnt)
	lines[2] = fmt.Sprint(g.width, g.height)

	text.Draw(screen, strings.Join(lines, "\n"), basicfont.Face7x13, 10, 20, color.White)
	vector.StrokeRect(screen, 1, 1, float32(g.width)-1, float32(g.height)-1, 1, color.White, false)
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	g.width, g.height = outsideWidth/2, outsideHeight/2
	return g.width, g.height
}
