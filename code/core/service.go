package core

type CalicoCutService interface {
	Get(countryCode string) (*Country, error)
}
