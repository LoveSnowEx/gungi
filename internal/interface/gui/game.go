package gui

import (
	"fmt"
	"time"

	"github.com/LoveSnowEx/gungi/internal/infra/persist"
	"github.com/LoveSnowEx/gungi/pkg/gungi/app/usecase"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var cnt int

type game struct {
	width, height int
	gungiUseCase  usecase.GameUsecase
	updateCh      chan struct{}
}

func newGame() *game {
	return &game{
		width:  320,
		height: 240,
		gungiUseCase: usecase.NewGameUsecase(
			&usecase.GameUsecaseConfig{
				GameRepo:   persist.NewGameRepo(),
				PlayerRepo: persist.NewPlayerRepo(),
			},
		),
		updateCh: make(chan struct{}, 64),
	}
}

func (g *game) Update() error {
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	select {
	case <-g.updateCh:
		cnt++
	default:
	}
	ebitenutil.DebugPrint(screen, fmt.Sprintf("Hello, World! %d", cnt))
}

func (g *game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}

func Run() (err error) {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	game := newGame()
	go func() {
		for {
			time.Sleep(1000 * time.Millisecond)
			game.updateCh <- struct{}{}
		}
	}()

	return ebiten.RunGame(game)
}
