package core

// CalicoCutService represents functions to be used by user-interfaces.
type CalicoCutService interface {
	Get(countryCode string) (*Country, error)
}
