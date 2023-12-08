package config

import "github.com/LoveSnowEx/gungi/pkg/gungi/domain/model"

var (
	PieceAmounts = map[model.PieceType]int{
		model.Marshal:           1,
		model.General:           1,
		model.LieutenantGeneral: 1,
		model.MajorGeneral:      2,
		model.Samurai:           2,
		model.Lancer:            3,
		model.Knight:            2,
		model.Spy:               2,
		model.Fortress:          2,
		model.Pawn:              4,
		model.Cannon:            1,
		model.Musketeer:         1,
		model.Archer:            2,
		model.Captain:           1,
	}
)
