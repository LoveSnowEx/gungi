package gungi_usecase

import (
	"github.com/LoveSnowEx/gungi/config"
	"github.com/LoveSnowEx/gungi/internal/const/gungi_errors"
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_model"
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_repo"
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_service"

	"github.com/gookit/event"
)

var (
	_ GameUsecase = (*gameUsecase)(nil)
)

type GameUsecaseConfig struct {
	GameRepo     gungi_repo.GameRepo
	PlayerRepo   gungi_repo.PlayerRepo
	EventManager event.ManagerFace
}

type GameUsecase interface {
	// CreateGame creates a new game.
	CreateGame() (gungi_model.Game, error)
	// FindGame finds a game by id.
	FindGame(id uint) (gungi_model.Game, error)
	// JoinGame adds a player to the game.
	JoinGame(gameId uint, playerId uint, color gungi_model.Color) error
	// LeaveGame removes a player from the game.
	LeaveGame(gameId uint, playerId uint) error
	// StartGame starts the game.
	StartGame(gameId uint) error
}

type gameUsecase struct {
	gameService  gungi_service.GameService
	gameRepo     gungi_repo.GameRepo
	playerRepo   gungi_repo.PlayerRepo
	eventManager event.ManagerFace
}

func NewGameUsecase(config *GameUsecaseConfig) GameUsecase {
	return &gameUsecase{
		gameService:  gungi_service.NewGameService(),
		gameRepo:     config.GameRepo,
		playerRepo:   config.PlayerRepo,
		eventManager: config.EventManager,
	}
}

func (u *gameUsecase) CreateGame() (game gungi_model.Game, err error) {
	if u == nil || u.gameRepo == nil || u.gameService == nil || u.playerRepo == nil {
		err = gungi_errors.ErrInvalidService
		return
	}
	game = u.gameService.Create()
	if err = u.gameRepo.Save(game); err != nil {
		game = nil
		return
	}
	return
}

func (u *gameUsecase) FindGame(id uint) (gungi_model.Game, error) {
	return u.gameRepo.Find(id)
}

func (u *gameUsecase) JoinGame(gameId uint, playerId uint, color gungi_model.Color) (err error) {
	var game gungi_model.Game
	game, err = u.gameRepo.Find(gameId)
	if err != nil {
		return
	}
	var player gungi_model.Player
	if player, err = u.playerRepo.Find(playerId); err != nil {
		return
	}
	if err = u.gameService.Join(game, player, color); err != nil {
		return
	}
	if err = u.gameRepo.Save(game); err != nil {
		return
	}
	if err, _ = u.eventManager.Fire("game.update.playerjoin", nil); err != nil {
		return
	}
	return
}

func (u *gameUsecase) LeaveGame(gameId uint, playerId uint) (err error) {
	game, err := u.gameRepo.Find(gameId)
	if err != nil {
		return
	}
	var player gungi_model.Player
	if player, err = u.playerRepo.Find(playerId); err != nil {
		return
	}
	if err = u.gameService.Leave(game, player); err != nil {
		return
	}
	if err = u.gameRepo.Save(game); err != nil {
		return
	}
	if err, _ = u.eventManager.Fire("game.update.playerleave", nil); err != nil {
		return
	}
	return
}

func (u *gameUsecase) StartGame(gameId uint) (err error) {
	var game gungi_model.Game
	if game, err = u.gameRepo.Find(gameId); err != nil {
		return
	}
	if err = u.gameService.Start(game, config.PieceAmounts); err != nil {
		return
	}
	if err = u.gameRepo.Save(game); err != nil {
		return
	}
	if err, _ = u.eventManager.Fire("game.update.start", nil); err != nil {
		return
	}
	return
}
