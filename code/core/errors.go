package core

const (
	UserErrPrefix      = "USERERR"
	UnknownErrorPrefix = "UNKNOWNERR"
)

const (
	CountryNotFoundErr string = UserErrPrefix + "-1"

	UnknownErr string = UnknownErrorPrefix + "-1"
)

var (
	CodeMapping = map[string]string{
		CountryNotFoundErr: "could not find the specified country",
		UnknownErr:         "unknown error",
	}
)

// Error is a custom type that adheres to the default Golang error interface. This allows us to store error codes that
// other parts of the code can reference for smarter reporting.
type Error struct {
	Code    string
	Message string
}

// Error returns the error message stored in the error. If no message is present, it attempts to return the default
// message for the error type.
func (e Error) Error() string {
	if e.Message != "" {
		return e.Message
	}

	if msg, ok := CodeMapping[e.Code]; ok {
		return msg
	}

	msg, _ := CodeMapping[UnknownErr]

	return msg
}
