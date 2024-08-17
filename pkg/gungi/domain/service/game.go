package service

import (
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"
	"github.com/LoveSnowEx/gungi/pkg/gungi/errors"
)

type GameService interface {
	// Create creates a new game.
	Create() model.Game
	// Join adds a player to the game.
	Join(game model.Game, player model.Player, team model.Color) error
	// Leave removes a player from the game.
	Leave(game model.Game, player model.Player) error
	// Start starts the game.
	Start(game model.Game, pieceAmounts map[model.PieceType]int) error
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

func (g *gameService) Start(game model.Game, pieceAmounts map[model.PieceType]int) (err error) {
	if game.Phase() != model.Setup {
		err = errors.ErrInvalidPhase
		return
	}
	if game.Player(model.White) == nil || game.Player(model.Black) == nil {
		err = errors.ErrInvalidPlayerAmount
		return
	}
	pieceId := uint(0)
	for _, color := range model.Colors() {
		for pieceType, amount := range pieceAmounts {
			for i := 0; i < amount; i++ {
				piece := model.NewPiece(pieceId, pieceType, color)
				if err = game.Reserve(color).Add(piece); err != nil {
					return
				}
				pieceId++
			}
		}
	}
	game.SetPhase(model.Prepare)
	return
}
