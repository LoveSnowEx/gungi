package service

import (
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/repo"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/service"
	"github.com/LoveSnowEx/gungi/pkg/gungi/infra/config"
	"github.com/LoveSnowEx/gungi/pkg/gungi/infra/persist"
)

type GameService interface {
	// CreateGame creates a new game.
	CreateGame() (model.Game, error)
	// FindGame finds a game by id.
	FindGame(id uint) (model.Game, error)
	// JoinGame adds a player to the game.
	JoinGame(gameId uint, playerId uint, color model.Color) error
	// LeaveGame removes a player from the game.
	LeaveGame(gameId uint, playerId uint) error
	// StartGame starts the game.
	StartGame(gameId uint) error
}

type gameService struct {
	gameRepo    repo.GameRepo
	gameService service.GameService
	playerRepo  repo.PlayerRepo
}

func NewGameService() GameService {
	return &gameService{
		gameRepo:    persist.NewGameRepo(),
		gameService: service.NewGameService(),
		playerRepo:  persist.NewPlayerRepo(),
	}
}

func (c *gameService) CreateGame() (game model.Game, err error) {
	game = c.gameService.Create()
	if err = c.gameRepo.Save(game); err != nil {
		game = nil
		return
	}
	return
}

func (c *gameService) FindGame(id uint) (model.Game, error) {
	return c.gameRepo.Find(id)
}

func (c *gameService) JoinGame(gameId uint, playerId uint, color model.Color) (err error) {
	var game model.Game
	game, err = c.gameRepo.Find(gameId)
	if err != nil {
		return
	}
	var player model.Player
	if player, err = c.playerRepo.Find(playerId); err != nil {
		return
	}
	if err = c.gameService.Join(game, player, color); err != nil {
		return
	}
	err = c.gameRepo.Save(game)
	return
}

func (c *gameService) LeaveGame(gameId uint, playerId uint) (err error) {
	game, err := c.gameRepo.Find(gameId)
	if err != nil {
		return
	}
	var player model.Player
	if player, err = c.playerRepo.Find(playerId); err != nil {
		return
	}
	if err = c.gameService.Leave(game, player); err != nil {
		return
	}
	err = c.gameRepo.Save(game)
	return
}

func (c *gameService) StartGame(gameId uint) (err error) {
	var game model.Game
	if game, err = c.gameRepo.Find(gameId); err != nil {
		return
	}
	if err = c.gameService.Start(game, config.PieceAmounts); err != nil {
		return
	}
	err = c.gameRepo.Save(game)
	return
}
