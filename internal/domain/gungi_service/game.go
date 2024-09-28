package gungi_service

import (
	"github.com/LoveSnowEx/gungi/internal/const/gungi_errors"
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_model"
)

type GameService interface {
	// Create creates a new game.
	Create() gungi_model.Game
	// Join adds a player to the game.
	Join(game gungi_model.Game, player gungi_model.Player, team gungi_model.Color) error
	// Leave removes a player from the game.
	Leave(game gungi_model.Game, player gungi_model.Player) error
	// Start starts the game.
	Start(game gungi_model.Game, pieceAmounts map[gungi_model.PieceType]int) error
}

type gameService struct {
}

func NewGameService() *gameService {
	return &gameService{}
}

func (g *gameService) Create() gungi_model.Game {
	return gungi_model.NewGame()
}

func (g *gameService) Join(game gungi_model.Game, player gungi_model.Player, color gungi_model.Color) error {
	return game.Join(color, player)
}

func (g *gameService) Leave(game gungi_model.Game, player gungi_model.Player) error {
	return game.Leave(player)
}

func (g *gameService) Start(game gungi_model.Game, pieceAmounts map[gungi_model.PieceType]int) (err error) {
	if game.Phase() != gungi_model.Setup {
		err = gungi_errors.ErrInvalidPhase
		return
	}
	if game.Player(gungi_model.White) == nil || game.Player(gungi_model.Black) == nil {
		err = gungi_errors.ErrInvalidPlayerAmount
		return
	}
	pieceId := uint(0)
	for _, color := range gungi_model.Colors() {
		for pieceType, amount := range pieceAmounts {
			for i := 0; i < amount; i++ {
				piece := gungi_model.NewPiece(pieceId, pieceType, color)
				if err = game.Reserve(color).Add(piece); err != nil {
					return
				}
				pieceId++
			}
		}
	}
	game.SetPhase(gungi_model.Prepare)
	return
}
