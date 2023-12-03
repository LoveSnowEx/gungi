package service

import (
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/repo"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/service"
)

type GameServiceConfig struct {
	GameRepository repo.GameRepo
	GameService    service.GameService
}

type GameService interface {
	// CreateGame creates a new game.
	CreateGame() error
	// JoinGame adds a player to the game.
	JoinGame(gameID uint, playerID uint, color string) error
	// LeaveGame removes a player from the game.
	LeaveGame(gameID uint, playerID uint) error
	// StartGame starts the game.
	StartGame(gameID uint) error
}

type gameService struct {
	gameRepository repo.GameRepo
	gameService    service.GameService
}

// func NewGameController(config GameServiceConfig) GameService {
// 	return &gameService{
// 		gameRepository: config.GameRepository,
// 		gameService:    config.GameService,
// 	}
// }

// func (c *gameService) CreateGame() (err error) {

// 	c.gameService.Create()
// }
