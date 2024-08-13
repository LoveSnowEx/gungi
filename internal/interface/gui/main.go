package gui

import "github.com/hajimehoshi/ebiten/v2"

const (
	windowWidth  = 1280
	windowHeight = 720
	tps          = 1
	title        = "Gungi"
)

func Run() (err error) {
	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetTPS(tps)
	ebiten.SetScreenClearedEveryFrame(true)
	ebiten.SetWindowTitle(title)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	game := newGame()
	return ebiten.RunGame(game)
}
