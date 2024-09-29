package persist

import (
	"context"
	"log/slog"

	"github.com/LoveSnowEx/gungi/internal/const/gungi_errors"
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_model"
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_repo"

	"github.com/LoveSnowEx/gungi/internal/infra/database"
	"github.com/LoveSnowEx/gungi/internal/infra/po"
	"gorm.io/gorm"
)

const boardSize = gungi_model.BoardRows * gungi_model.BoardCols * gungi_model.BoardLevels

var (
	_ gungi_repo.GameRepo = (*gameRepoImpl)(nil)
)

type gameRepoImpl struct {
	playerRepo gungi_repo.PlayerRepo
}

func NewGameRepo() gungi_repo.GameRepo {
	return &gameRepoImpl{
		playerRepo: NewPlayerRepo(),
	}
}

func (r *gameRepoImpl) Find(id uint) (game gungi_model.Game, err error) {
	gamePo := po.Game{}
	err = database.Default().
		WithContext(context.Background()).
		Where("id = ?", id).
		Preload("Players").
		Preload("Players.User").
		Preload("Pieces").
		First(&gamePo).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			slog.Warn("game not found", "id", id)
			err = gungi_errors.ErrGameNotFound
		}
		return
	}
	game = gungi_model.NewGame()
	game.SetId(gamePo.ID)
	game.SetTurn(gamePo.Turn)
	game.SetPhase(gamePo.Phase)
	// Players
	for _, playerPo := range gamePo.Players {
		user := playerPo.User
		player := gungi_model.NewPlayer(user.Name)
		player.SetId(user.ID)
		err = game.Join(playerPo.Color, player)
		if err != nil {
			game = nil
			return
		}
	}
	board := gungi_model.NewBoard()
	reserve := map[gungi_model.Color]gungi_model.PieceArea{}
	discard := map[gungi_model.Color]gungi_model.PieceArea{}

p:
	for _, piece := range gamePo.Pieces {
		p := gungi_model.NewPiece(piece.Type, piece.Color)
		position := piece.Position
		if position < uint(boardSize) {
			z := position % gungi_model.BoardLevels
			y := (position / gungi_model.BoardLevels) % gungi_model.BoardCols
			x := position / (gungi_model.BoardCols * gungi_model.BoardLevels)
			board.Set(gungi_model.NewBoardPosition(x, y, z), p)
			continue p
		}
		position -= uint(boardSize)
		for _, color := range gungi_model.Colors() {
			reserve[color] = gungi_model.NewPieceArea()
			if position < gungi_model.AreaSize {
				reserve[color].Set(position, p)
				continue p
			}
			position -= gungi_model.AreaSize
		}
		for _, color := range gungi_model.Colors() {
			discard[color] = gungi_model.NewPieceArea()
			if position < gungi_model.AreaSize {
				discard[color].Set(position, p)
				continue p
			}
			position -= gungi_model.AreaSize
		}
	}
	return
}

func (r *gameRepoImpl) Save(game gungi_model.Game) (err error) {
	if game == nil {
		err = gungi_errors.ErrInvalidGame
		return
	}
	gamePo := po.Game{
		Players: make([]po.Player, 0),
		Pieces:  make([]po.Piece, 0),
		Turn:    game.Turn(),
		Phase:   game.Phase(),
	}

	// Setup ID
	if game.Id() != 0 {
		gamePo.ID = game.Id()
	} else {
		if err = database.Default().Save(&gamePo).Error; err != nil {
			return
		}
	}
	// Board
	for x := range [gungi_model.BoardRows]struct{}{} {
		for y := range [gungi_model.BoardCols]struct{}{} {
			for z := range [gungi_model.BoardLevels]struct{}{} {
				if piece := game.Board()[x][y][z]; piece != nil {
					gamePo.Pieces = append(gamePo.Pieces, po.Piece{
						GameID:   gamePo.ID,
						Type:     piece.Type(),
						Color:    piece.Color(),
						Position: uint(((x*gungi_model.BoardCols)+y)*gungi_model.BoardLevels + z),
					})
				}
			}
		}
	}
	// Players, Reserve, Discard
	for _, color := range gungi_model.Colors() {
		// Players
		if player := game.Player(color); player != nil {
			if player.Id() == 0 {
				return gungi_errors.ErrInvalidPlayer
			}
			playerPo := po.Player{
				UserID: player.Id(),
				GameID: gamePo.ID,
				Color:  color,
			}
			gamePo.Players = append(gamePo.Players, playerPo)
		}
		// Reserve
		for i, piece := range game.Reserve(color) {
			if piece == nil {
				continue
			}
			position := uint(boardSize+gungi_model.AreaSize*color) + uint(i)
			gamePo.Pieces = append(gamePo.Pieces, po.Piece{
				GameID:   gamePo.ID,
				Type:     piece.Type(),
				Color:    piece.Color(),
				Position: position,
			})
		}
		// Discard
		for i, piece := range game.Discard(color) {
			if piece == nil {
				continue
			}
			position := uint(boardSize+gungi_model.AreaSize*2+gungi_model.AreaSize*color) + uint(i)
			gamePo.Pieces = append(gamePo.Pieces, po.Piece{
				GameID:   gamePo.ID,
				Type:     piece.Type(),
				Color:    piece.Color(),
				Position: position,
			})
		}
	}
	if err = database.Default().Save(&gamePo).Error; err != nil {
		return
	}
	game.SetId(gamePo.ID)
	return
}
