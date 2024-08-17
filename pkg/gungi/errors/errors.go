package errors

type GungiError uint

const (
	ErrInvalidGame GungiError = iota
	ErrInvalidPieceType
	ErrInvalidPhase
	ErrInvalidPlayerAmount
	ErrInvalidService
	ErrInvalidPlayer
	ErrPlayerAlreadyJoined
	ErrPlayerNotFound
	ErrTeamFull
	ErrPieceAlreadyExists
	ErrPieceNotFound
	ErrGameNotFound
)

var gungiErrors = map[GungiError]string{
	ErrInvalidGame:         "invalid game",
	ErrInvalidPieceType:    "invalid piece type",
	ErrInvalidPhase:        "invalid phase",
	ErrInvalidPlayerAmount: "invalid player amount",
	ErrInvalidService:      "invalid service",
	ErrInvalidPlayer:       "invalid player",
	ErrPlayerAlreadyJoined: "player already joined",
	ErrPlayerNotFound:      "player not found",
	ErrTeamFull:            "team full",
	ErrPieceAlreadyExists:  "piece already exists",
	ErrPieceNotFound:       "piece not found",
	ErrGameNotFound:        "game not found",
}

func (e GungiError) Error() string {
	if msg, ok := gungiErrors[e]; ok {
		return msg
	}
	return "unknown error"
}
