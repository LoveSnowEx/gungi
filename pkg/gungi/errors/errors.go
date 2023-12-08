package errors

type GungiError uint64

const (
	ErrInvalidGame GungiError = iota
	ErrInvalidPieceType
	ErrInvalidPhase
	ErrInvalidPlayerAmount
	ErrInvalidService
	ErrPlayerAlreadyJoined
	ErrPlayerNotFound
	ErrTeamFull
	ErrPieceAlreadyExists
	ErrPieceNotFound
	ErrGameNotFound
)

func (e GungiError) Error() string {
	return "unknown error"
}
