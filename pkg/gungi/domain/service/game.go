package service

import (
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"
)

type GameService interface {
	// Create creates a new game.
	Create() model.Game
	// Join adds a player to the game.
	Join(game model.Game, player model.Player, team model.Color) error
	// Leave removes a player from the game.
	Leave(game model.Game, player model.Player) error
	// Start starts the game.
	Start(game model.Game) error
}

type gameService struct {
}

func NewGameService() GameService {
	return &gameService{}
}

func (g *gameService) Create() model.Game {
	return model.NewGame()
}

func (g *gameService) Join(game model.Game, player model.Player, color model.Color) error {
	return game.Join(color, player)
}

func (g *gameService) Leave(game model.Game, player model.Player) error {
	return game.Leave(player)
}

func (g *gameService) Start(game model.Game) (err error) {
	return
}
