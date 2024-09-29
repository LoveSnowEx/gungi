package po

import (
	"github.com/LoveSnowEx/gungi/internal/domain/gungi_model"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Players []Player
	Pieces  []Piece
	Turn    gungi_model.Color
	Phase   gungi_model.Phase
	Winner  gungi_model.Color
}

type Piece struct {
	gorm.Model
	GameID   uint `gorm:"primaryKey"`
	Type     gungi_model.PieceType
	Color    gungi_model.Color
	Position uint
}

type Player struct {
	GameID uint `gorm:"primaryKey"`
	UserID uint `gorm:"primaryKey"`
	User   User
	Color  gungi_model.Color
}
