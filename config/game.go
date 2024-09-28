package config

import "github.com/LoveSnowEx/gungi/internal/domain/gungi_model"

var (
	PieceAmounts = map[gungi_model.PieceType]int{
		gungi_model.Marshal:           1,
		gungi_model.General:           1,
		gungi_model.LieutenantGeneral: 1,
		gungi_model.MajorGeneral:      2,
		gungi_model.Samurai:           2,
		gungi_model.Lancer:            3,
		gungi_model.Knight:            2,
		gungi_model.Spy:               2,
		gungi_model.Fortress:          2,
		gungi_model.Pawn:              4,
		gungi_model.Cannon:            1,
		gungi_model.Musketeer:         1,
		gungi_model.Archer:            2,
		gungi_model.Captain:           1,
	}
)
