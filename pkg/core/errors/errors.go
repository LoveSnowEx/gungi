package errors

type CoreError uint

const (
	ErrUserNotFound CoreError = iota
)

var coreErrors = map[CoreError]string{
	ErrUserNotFound: "user not found",
}

func (e CoreError) Error() string {
	if msg, ok := coreErrors[e]; ok {
		return msg
	}
	return "unknown error"
}
