package po

import (
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
	Color Color
}

func ToColor(color Color) model.Color {
	switch color {
	case White:
		return model.White
	case Black:
		return model.Black
	}
	panic("invalid color")
}

func FromColor(color model.Color) Color {
	switch color {
	case model.White:
		return White
	case model.Black:
		return Black
	}
	panic("invalid color")
}

func ToPhase(phase Phase) model.Phase {
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

func FromPhase(phase model.Phase) Phase {
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

func ToPieceType(pieceType PieceType) model.PieceType {
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

func FromPieceType(pieceType model.PieceType) PieceType {
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
