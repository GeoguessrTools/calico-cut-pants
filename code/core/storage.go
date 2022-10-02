package core

type Storage interface {
	GetCountry(country string) (*Country, error)
	GetAllCountries() ([]Country, error)
}
