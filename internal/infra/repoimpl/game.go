package persist

import (
	"encoding/json"

	"github.com/LoveSnowEx/gungi/internal/infra/po"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"
	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/repo"
	"github.com/LoveSnowEx/gungi/pkg/gungi/errors"
)

var (
	_           repo.GameRepo = (*gameRepoImpl)(nil)
	gameStorage               = make(map[uint][]byte)
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
	jsonBytes, ok := gameStorage[id]
	if !ok {
		err = errors.ErrGameNotFound
		return
	}
	gamePo := po.Game{}
	if err = json.Unmarshal(jsonBytes, &gamePo); err != nil {
		return
	}
	players := make([]model.Player, 0, len(gamePo.Players))
	for _, playerPo := range gamePo.Players {
		player, err := r.playerRepo.Find(playerPo.Id)
		if err != nil {
			return nil, err
		}
		players = append(players, player)
	}
	game = model.NewGame()
	game.SetId(gamePo.Id)
	game.SetCurrentTurn(po.ToColor(gamePo.CurrentTurn))
	game.SetPhase(po.ToPhase(gamePo.Phase))
	// Board
	for _, piecePo := range gamePo.BoardPieces {
		piece := model.NewPiece(piecePo.Id, po.ToPieceType(piecePo.Type), po.ToColor(piecePo.Color))
		game.Board().Set(model.NewVector3D(piecePo.Row, piecePo.Column, 0), piece)
	}
	// Reserve
	for _, piecePo := range gamePo.Reserve {
		piece := model.NewPiece(piecePo.Id, po.ToPieceType(piecePo.Type), po.ToColor(piecePo.Color))
		if err = game.Reserve(po.ToColor(piecePo.Color)).Add(piece); err != nil {
			game = nil
			return
		}
	}
	// Discard
	for _, piecePo := range gamePo.Discard {
		piece := model.NewPiece(piecePo.Id, po.ToPieceType(piecePo.Type), po.ToColor(piecePo.Color))
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
	gamePo := po.Game{
		Id: game.Id(),
	}
	for _, color := range model.Colors() {
		if player := game.Player(color); player != nil {
			gamePo.Players = append(gamePo.Players, po.Player{
				Id:    player.Id(),
				Color: po.FromColor(color),
			})
		}
		// Reserve
		for _, piece := range game.Reserve(color).Pieces() {
			piecePo := po.Piece{
				Id:    piece.Id(),
				Type:  po.FromPieceType(piece.Type()),
				Color: po.FromColor(piece.Color()),
			}
			gamePo.Reserve = append(gamePo.Reserve, piecePo)
		}
		// Discard
		for _, piece := range game.Discard(color).Pieces() {
			piecePo := po.Piece{
				Id:    piece.Id(),
				Type:  po.FromPieceType(piece.Type()),
				Color: po.FromColor(piece.Color()),
			}
			gamePo.Discard = append(gamePo.Discard, piecePo)
		}
	}
	if gamePo.Id == 0 {
		gamePo.Id = uint(len(gameStorage) + 1)
	}
	jsonBytes, err := json.Marshal(gamePo)
	if err != nil {
		return
	}
	gameStorage[game.Id()] = jsonBytes
	return
}
