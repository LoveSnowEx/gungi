package po

import (
	"encoding/json"

	"github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"
)

type Color uint

const (
	White Color = iota
	Black
)

type Phase uint

const (
	Setup Phase = iota
	Prepare
	Play
	End
)

type PieceType uint

const (
	Marshal           PieceType = iota // 帥
	General                            // 大
	LieutenantGeneral                  // 中
	MajorGeneral                       // 小
	Samurai                            // 侍
	Lancer                             // 槍
	Knight                             // 馬
	Spy                                // 忍
	Fortress                           // 砦
	Pawn                               // 兵
	Cannon                             // 砲
	Musketeer                          // 筒
	Archer                             // 弓
	Captain                            // 謀
)

type Game struct {
	Id          uint
	Players     []Player
	BoardPieces []BoardPiece
	Reserve     []Piece
	Discard     []Piece
	CurrentTurn Color
	Phase       Phase
}

type Piece struct {
	Id    uint
	Type  PieceType
	Color Color
}

type BoardPiece struct {
	Piece
	Row    int
	Column int
}

type Player struct {
	Id    uint
	Name  string
	Color Color
}

type GameMarshaller interface {
	Marshal(game model.Game) ([]byte, error)
	Unmarshal([]byte) (model.Game, error)
}

type gameMarshaller struct {
}

func NewGameMarshaller() GameMarshaller {
	return &gameMarshaller{}
}

func (m *gameMarshaller) Marshal(game model.Game) ([]byte, error) {
	gamePo := Game{}
	gamePo.Id = game.Id()
	gamePo.CurrentTurn = fromColor(game.CurrentTurn())
	gamePo.Phase = fromPhase(game.Phase())
	gamePo.BoardPieces = make([]BoardPiece, 0)
	gamePo.Reserve = make([]Piece, 0)
	gamePo.Discard = make([]Piece, 0)
	// Board
	for row := range [model.BoardRows]struct{}{} {
		for col := range [model.BoardCols]struct{}{} {
			for level := range [model.BoardLevels]struct{}{} {
				piece := game.Board().Get(model.NewVector3D(row, col, level))
				if piece == nil {
					continue
				}
				piecePo := BoardPiece{}
				piecePo.Id = piece.Id()
				piecePo.Type = fromPieceType(piece.Type())
				piecePo.Color = fromColor(piece.Color())
				piecePo.Row = row
				piecePo.Column = col
				gamePo.BoardPieces = append(gamePo.BoardPieces, piecePo)
			}
		}
	}
	// Player, Reserve, Discard
	for _, color := range model.Colors() {
		// Player
		player := game.Player(color)
		playerPo := &Player{}
		playerPo.Id = player.Id()
		playerPo.Name = player.Name()
		playerPo.Color = fromColor(color)
		// Reserve
		for _, piece := range game.Reserve(color).Pieces() {
			piecePo := Piece{}
			piecePo.Id = piece.Id()
			piecePo.Type = fromPieceType(piece.Type())
			piecePo.Color = fromColor(piece.Color())
			gamePo.Reserve = append(gamePo.Reserve, piecePo)
		}
		// Discard
		for _, piece := range game.Discard(color).Pieces() {
			piecePo := Piece{}
			piecePo.Id = piece.Id()
			piecePo.Type = fromPieceType(piece.Type())
			piecePo.Color = fromColor(piece.Color())
			gamePo.Discard = append(gamePo.Discard, piecePo)
		}
	}
	return json.Marshal(gamePo)
}

func (m *gameMarshaller) Unmarshal(bytes []byte) (game model.Game, err error) {
	gamePo := Game{}
	if err := json.Unmarshal(bytes, &gamePo); err != nil {
		return nil, err
	}
	game = model.NewGame()
	game.SetId(gamePo.Id)
	game.SetCurrentTurn(toColor(gamePo.CurrentTurn))
	game.SetPhase(toPhase(gamePo.Phase))
	// Board
	for _, piecePo := range gamePo.BoardPieces {
		piece := model.NewPiece(piecePo.Id, toPieceType(piecePo.Type), toColor(piecePo.Color))
		game.Board().Set(model.NewVector3D(piecePo.Row, piecePo.Column, 0), piece)
	}
	// Player, Reserve, Discard
	for _, playerPo := range gamePo.Players {
		// Player
		player := model.NewHumanPlayer(playerPo.Name)
		player.SetId(playerPo.Id)
		if err = game.Join(toColor(playerPo.Color), player); err != nil {
			game = nil
			return
		}
		// Reserve
		for _, piecePo := range gamePo.Reserve {
			piece := model.NewPiece(piecePo.Id, toPieceType(piecePo.Type), toColor(piecePo.Color))
			if err = game.Reserve(toColor(piecePo.Color)).Add(piece); err != nil {
				game = nil
				return
			}
		}
		// Discard
		for _, piecePo := range gamePo.Discard {
			piece := model.NewPiece(piecePo.Id, toPieceType(piecePo.Type), toColor(piecePo.Color))
			if err = game.Discard(toColor(piecePo.Color)).Add(piece); err != nil {
				game = nil
				return
			}
		}
	}
	return game, nil
}

func toColor(color Color) model.Color {
	switch color {
	case White:
		return model.White
	case Black:
		return model.Black
	}
	panic("invalid color")
}

func fromColor(color model.Color) Color {
	switch color {
	case model.White:
		return White
	case model.Black:
		return Black
	}
	panic("invalid color")
}

func toPhase(phase Phase) model.Phase {
	switch phase {
	case Setup:
		return model.Setup
	case Prepare:
		return model.Prepare
	case Play:
		return model.Play
	case End:
		return model.End
	}
	panic("invalid phase")
}

func fromPhase(phase model.Phase) Phase {
	switch phase {
	case model.Setup:
		return Setup
	case model.Prepare:
		return Prepare
	case model.Play:
		return Play
	case model.End:
		return End
	}
	panic("invalid phase")
}

func toPieceType(pieceType PieceType) model.PieceType {
	switch pieceType {
	case Marshal:
		return model.Marshal
	case General:
		return model.General
	case LieutenantGeneral:
		return model.LieutenantGeneral
	case MajorGeneral:
		return model.MajorGeneral
	case Samurai:
		return model.Samurai
	case Lancer:
		return model.Lancer
	case Knight:
		return model.Knight
	case Spy:
		return model.Spy
	case Fortress:
		return model.Fortress
	case Pawn:
		return model.Pawn
	case Cannon:
		return model.Cannon
	case Musketeer:
		return model.Musketeer
	case Archer:
		return model.Archer
	case Captain:
		return model.Captain
	}
	panic("invalid piece type")
}

func fromPieceType(pieceType model.PieceType) PieceType {
	switch pieceType {
	case model.Marshal:
		return Marshal
	case model.General:
		return General
	case model.LieutenantGeneral:
		return LieutenantGeneral
	case model.MajorGeneral:
		return MajorGeneral
	case model.Samurai:
		return Samurai
	case model.Lancer:
		return Lancer
	case model.Knight:
		return Knight
	case model.Spy:
		return Spy
	case model.Fortress:
		return Fortress
	case model.Pawn:
		return Pawn
	case model.Cannon:
		return Cannon
	case model.Musketeer:
		return Musketeer
	case model.Archer:
		return Archer
	case model.Captain:
		return Captain
	}
	panic("invalid piece type")
}
