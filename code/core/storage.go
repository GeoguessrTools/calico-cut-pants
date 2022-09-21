package core

type Storage interface {
	Get(country string) (*Country, error)
}
