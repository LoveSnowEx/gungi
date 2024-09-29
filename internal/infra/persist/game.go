package persist

import (
	"context"
	"log/slog"

	"github.com/LoveSnowEx/gungi/internal/const/gungi_errors"
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_model"
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_repo"
	"github.com/LoveSnowEx/gungi/internal/infra/dal"
	"github.com/LoveSnowEx/gungi/internal/infra/po"
	"gorm.io/gorm"
)

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
			err = gungi_errors.ErrGameNotFound
		}
		return
	}
	game = gungi_model.NewGame()
	game.SetId(gamePo.ID)
	game.SetTurn(po.ToColor(gamePo.CurrentTurn))
	game.SetPhase(po.ToPhase(gamePo.Phase))
	// Players
	for _, playerPo := range gamePo.Players {
		user := playerPo.User
		player := gungi_model.NewPlayer(user.Name)
		player.SetId(user.ID)
		err = game.Join(po.ToColor(playerPo.Color), player)
		if err != nil {
			game = nil
			return
		}
	}
	// Board
	for _, piecePo := range gamePo.BoardPieces {
		piece := gungi_model.NewPiece(piecePo.PieceID, po.ToPieceType(piecePo.Type), po.ToColor(piecePo.Color))
		game.Board().Set(gungi_model.NewVector3D(piecePo.Row, piecePo.Column, 0), piece)
	}
	// Reserve
	for _, piecePo := range gamePo.Reserve {
		piece := gungi_model.NewPiece(piecePo.PieceID, po.ToPieceType(piecePo.Type), po.ToColor(piecePo.Color))
		if err = game.Reserve(po.ToColor(piecePo.Color)).Add(piece); err != nil {
			game = nil
			return
		}
	}
	// Discard
	for _, piecePo := range gamePo.Discard {
		piece := gungi_model.NewPiece(piecePo.PieceID, po.ToPieceType(piecePo.Type), po.ToColor(piecePo.Color))
		if err = game.Discard(po.ToColor(piecePo.Color)).Add(piece); err != nil {
			game = nil
			return
		}
	}
	return
}

func (r *gameRepoImpl) Save(game gungi_model.Game) (err error) {
	if game == nil {
		err = gungi_errors.ErrInvalidGame
		return
	}
	g := dal.Game
	gamePo := po.Game{
		Players:     make([]po.Player, 0),
		BoardPieces: make([]po.BoardPiece, 0),
		Reserve:     make([]po.ReservePiece, 0),
		Discard:     make([]po.DiscardPiece, 0),
		CurrentTurn: po.FromColor(game.Turn()),
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
	for x := range [gungi_model.BoardRows]struct{}{} {
		for y := range [gungi_model.BoardCols]struct{}{} {
			for z := range [gungi_model.BoardLevels]struct{}{} {
				if piece := game.Board().Get(gungi_model.NewVector3D(x, y, z)); piece != nil {
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
	for _, color := range gungi_model.Colors() {
		// Players
		if player := game.Player(color); player != nil {
			if player.Id() == 0 {
				return gungi_errors.ErrInvalidPlayer
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
			piecePo := po.ReservePiece{
				Piece: po.Piece{
					PieceID: piece.Id(),
					GameID:  gamePo.ID,
					Type:    po.FromPieceType(piece.Type()),
					Color:   po.FromColor(piece.Color()),
				},
			}
			gamePo.Reserve = append(gamePo.Reserve, piecePo)
		}
		// Discard
		for _, piece := range game.Discard(color).Pieces() {
			piecePo := po.DiscardPiece{
				Piece: po.Piece{
					PieceID: piece.Id(),
					GameID:  gamePo.ID,
					Type:    po.FromPieceType(piece.Type()),
					Color:   po.FromColor(piece.Color()),
				},
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
