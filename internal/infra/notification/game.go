package notification

import "github.com/gookit/event"

var (
	gameEventMgr *GameManager
)

type GameManager struct {
	*event.Manager
}

func NewGameManager() *GameManager {
	if gameEventMgr != nil {
		return gameEventMgr
	}
	gameEventMgr = &GameManager{
		event.NewManager("game", event.UsePathMode),
	}
	return gameEventMgr
}
