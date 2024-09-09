package po

import (
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_model"
	"gorm.io/gorm"
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
	gorm.Model
	Players     []Player
	BoardPieces []BoardPiece
	Reserve     []ReservePiece
	Discard     []DiscardPiece
	CurrentTurn Color
	Phase       Phase
}

type Piece struct {
	GameID  uint `gorm:"primaryKey"`
	PieceID uint `gorm:"primaryKey"`
	Type    PieceType
	Color   Color
}

type BoardPiece struct {
	Piece  `gorm:"embedded"`
	Row    int
	Column int
	Level  int
}

type ReservePiece struct {
	Piece `gorm:"embedded"`
}

type DiscardPiece struct {
	Piece `gorm:"embedded"`
}

type Player struct {
	GameID uint `gorm:"primaryKey"`
	UserID uint `gorm:"primaryKey"`
	User   User
	Color  Color
}

func ToColor(color Color) gungi_model.Color {
	switch color {
	case White:
		return gungi_model.White
	case Black:
		return gungi_model.Black
	}
	panic("invalid color")
}

func FromColor(color gungi_model.Color) Color {
	switch color {
	case gungi_model.White:
		return White
	case gungi_model.Black:
		return Black
	}
	panic("invalid color")
}

func ToPhase(phase Phase) gungi_model.Phase {
	switch phase {
	case Setup:
		return gungi_model.Setup
	case Prepare:
		return gungi_model.Prepare
	case Play:
		return gungi_model.Play
	case End:
		return gungi_model.End
	}
	panic("invalid phase")
}

func FromPhase(phase gungi_model.Phase) Phase {
	switch phase {
	case gungi_model.Setup:
		return Setup
	case gungi_model.Prepare:
		return Prepare
	case gungi_model.Play:
		return Play
	case gungi_model.End:
		return End
	}
	panic("invalid phase")
}

func ToPieceType(pieceType PieceType) gungi_model.PieceType {
	switch pieceType {
	case Marshal:
		return gungi_model.Marshal
	case General:
		return gungi_model.General
	case LieutenantGeneral:
		return gungi_model.LieutenantGeneral
	case MajorGeneral:
		return gungi_model.MajorGeneral
	case Samurai:
		return gungi_model.Samurai
	case Lancer:
		return gungi_model.Lancer
	case Knight:
		return gungi_model.Knight
	case Spy:
		return gungi_model.Spy
	case Fortress:
		return gungi_model.Fortress
	case Pawn:
		return gungi_model.Pawn
	case Cannon:
		return gungi_model.Cannon
	case Musketeer:
		return gungi_model.Musketeer
	case Archer:
		return gungi_model.Archer
	case Captain:
		return gungi_model.Captain
	}
	panic("invalid piece type")
}

func FromPieceType(pieceType gungi_model.PieceType) PieceType {
	switch pieceType {
	case gungi_model.Marshal:
		return Marshal
	case gungi_model.General:
		return General
	case gungi_model.LieutenantGeneral:
		return LieutenantGeneral
	case gungi_model.MajorGeneral:
		return MajorGeneral
	case gungi_model.Samurai:
		return Samurai
	case gungi_model.Lancer:
		return Lancer
	case gungi_model.Knight:
		return Knight
	case gungi_model.Spy:
		return Spy
	case gungi_model.Fortress:
		return Fortress
	case gungi_model.Pawn:
		return Pawn
	case gungi_model.Cannon:
		return Cannon
	case gungi_model.Musketeer:
		return Musketeer
	case gungi_model.Archer:
		return Archer
	case gungi_model.Captain:
		return Captain
	}
	panic("invalid piece type")
}
