package gui

import "github.com/hajimehoshi/ebiten/v2"

func Run() (err error) {
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetTPS(1)
	ebiten.SetScreenClearedEveryFrame(true)
	ebiten.SetWindowTitle("Hello, World!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	game := newGame()
	return ebiten.RunGame(game)
}
