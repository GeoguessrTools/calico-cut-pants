package core

type CalicoCutService interface {
	GetCountry(countryCode string) (*Country, error)
	GetAllCountryNames() ([]string, error)
}
