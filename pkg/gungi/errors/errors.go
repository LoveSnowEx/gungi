package errors

type GungiError uint64

const (
	ErrInvalidPieceType GungiError = iota
	ErrPlayerAlreadyJoined
	ErrPlayerNotFound
	ErrTeamFull
	ErrPieceAlreadyExists
	ErrPieceNotFound
	ErrGameNotFound
	ErrInvalidPhase
	ErrInvalidPlayerAmount
)

func (e GungiError) Error() string {
	return "unknown error"
}
