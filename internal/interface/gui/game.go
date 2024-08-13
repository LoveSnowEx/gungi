package gui

import (
	"github.com/LoveSnowEx/gungi/internal/infra/notification"
	"github.com/LoveSnowEx/gungi/internal/infra/persist"
	"github.com/LoveSnowEx/gungi/pkg/gungi/app/usecase"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"
	"github.com/gookit/event"
)

var cnt int

type game struct {
	width, height int
	gungiUseCase  usecase.GameUsecase
	gameInstance  *model.Game
	updateCh      chan struct{}
}

func newGame() (g *game) {
	g = &game{
		width:    160,
		height:   120,
		updateCh: make(chan struct{}, 64),
	}
	// Setup event manager
	gameEventManager := notification.NewGameManager()
	updateEventListener := event.ListenerFunc(func(_ event.Event) error {
		g.updateCh <- struct{}{}
		return nil
	})
	gameEventManager.AddListener("game.update.*", updateEventListener)
	// Setup usecase
	g.gungiUseCase = usecase.NewGameUsecase(
		&usecase.GameUsecaseConfig{
			GameRepo:     persist.NewGameRepo(),
			PlayerRepo:   persist.NewPlayerRepo(),
			EventManager: gameEventManager,
		},
	)
	return
}
