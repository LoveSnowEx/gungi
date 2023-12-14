package persist

import (
	"context"
	"log/slog"

	"github.com/LoveSnowEx/gungi/internal/infra/dal"
	"github.com/LoveSnowEx/gungi/internal/infra/po"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/repo"
	"github.com/LoveSnowEx/gungi/pkg/gungi/errors"
	"gorm.io/gorm"
)

var (
	_ repo.GameRepo = (*gameRepoImpl)(nil)
)

type gameRepoImpl struct {
	playerRepo repo.PlayerRepo
}

func NewGameRepo() repo.GameRepo {
	return &gameRepoImpl{
		playerRepo: NewPlayerRepo(),
	}
}

func (r *gameRepoImpl) Find(id uint) (game model.Game, err error) {
	g := dal.Game
	gamePo, err := g.
		WithContext(context.Background()).
		Where(g.ID.Eq(id)).
		Preload(g.Players).
		Preload(g.Players.User.RelationField).
		Preload(g.BoardPieces).
		Preload(g.Reserve).
		Preload(g.Discard).
		First()
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			slog.Warn("game not found", "id", id)
			err = errors.ErrGameNotFound
		}
		return
	}
	game = model.NewGame()
	game.SetId(gamePo.ID)
	game.SetCurrentTurn(po.ToColor(gamePo.CurrentTurn))
	game.SetPhase(po.ToPhase(gamePo.Phase))
	// Players
	for _, playerPo := range gamePo.Players {
		user := playerPo.User
		player := model.NewPlayer(user.Name)
		player.SetId(user.ID)
		err = game.Join(po.ToColor(playerPo.Color), player)
		if err != nil {
			game = nil
			return
		}
	}
	// Board
	for _, piecePo := range gamePo.BoardPieces {
		piece := model.NewPiece(piecePo.PieceID, po.ToPieceType(piecePo.Type), po.ToColor(piecePo.Color))
		game.Board().Set(model.NewVector3D(piecePo.Row, piecePo.Column, 0), piece)
	}
	// Reserve
	for _, piecePo := range gamePo.Reserve {
		piece := model.NewPiece(piecePo.PieceID, po.ToPieceType(piecePo.Type), po.ToColor(piecePo.Color))
		if err = game.Reserve(po.ToColor(piecePo.Color)).Add(piece); err != nil {
			game = nil
			return
		}
	}
	// Discard
	for _, piecePo := range gamePo.Discard {
		piece := model.NewPiece(piecePo.PieceID, po.ToPieceType(piecePo.Type), po.ToColor(piecePo.Color))
		if err = game.Discard(po.ToColor(piecePo.Color)).Add(piece); err != nil {
			game = nil
			return
		}
	}
	return
}

func (r *gameRepoImpl) Save(game model.Game) (err error) {
	if game == nil {
		err = errors.ErrInvalidGame
		return
	}
	g := dal.Game
	gamePo := po.Game{
		Players:     make([]po.Player, 0),
		BoardPieces: make([]po.BoardPiece, 0),
		Reserve:     make([]po.Piece, 0),
		Discard:     make([]po.Piece, 0),
		CurrentTurn: po.FromColor(game.CurrentTurn()),
		Phase:       po.FromPhase(game.Phase()),
	}

	// Setup ID
	if game.Id() != 0 {
		gamePo.ID = game.Id()
	} else {
		err = g.WithContext(context.Background()).Save(&gamePo)
		if err != nil {
			return
		}
	}
	// Board
	for x := range [model.BoardRows]struct{}{} {
		for y := range [model.BoardCols]struct{}{} {
			for z := range [model.BoardLevels]struct{}{} {
				if piece := game.Board().Get(model.NewVector3D(x, y, z)); piece != nil {
					piecePo := po.BoardPiece{
						Piece: po.Piece{
							PieceID: piece.Id(),
							GameID:  gamePo.ID,
							Type:    po.FromPieceType(piece.Type()),
							Color:   po.FromColor(piece.Color()),
						},
						Row:    x,
						Column: y,
						Level:  z,
					}
					gamePo.BoardPieces = append(gamePo.BoardPieces, piecePo)
				}
			}
		}
	}
	// Players, Reserve, Discard
	for _, color := range model.Colors() {
		// Players
		if player := game.Player(color); player != nil {
			if player.Id() == 0 {
				return errors.ErrInvalidPlayer
			}
			playerPo := po.Player{
				UserID: player.Id(),
				GameID: gamePo.ID,
				Color:  po.FromColor(color),
			}
			gamePo.Players = append(gamePo.Players, playerPo)
		}
		// Reserve
		for _, piece := range game.Reserve(color).Pieces() {
			piecePo := po.Piece{
				PieceID: piece.Id(),
				GameID:  gamePo.ID,
				Type:    po.FromPieceType(piece.Type()),
				Color:   po.FromColor(piece.Color()),
			}
			gamePo.Reserve = append(gamePo.Reserve, piecePo)
		}
		// Discard
		for _, piece := range game.Discard(color).Pieces() {
			piecePo := po.Piece{
				PieceID: piece.Id(),
				GameID:  gamePo.ID,
				Type:    po.FromPieceType(piece.Type()),
				Color:   po.FromColor(piece.Color()),
			}
			gamePo.Discard = append(gamePo.Discard, piecePo)
		}
	}
	err = g.WithContext(context.Background()).Save(&gamePo)
	if err != nil {
		return
	}
	game.SetId(gamePo.ID)
	return
}
