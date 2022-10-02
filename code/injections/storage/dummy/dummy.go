package dummy

import (
	"fmt"
	"github.com/GeoguessrTools/calico-cut-pants/core"
)

type dummyImpl struct {
}

func NewDummyStorage() (core.Storage, error) {
	return &dummyImpl{}, nil
}

func (d dummyImpl) GetCountry(countryId string) (*core.Country, error) {
	if country, ok := storedCountries[countryId]; ok {
		return &country, nil
	}

	return nil, fmt.Errorf("could not find country %s", countryId)
}

func (d dummyImpl) GetAllCountries() ([]core.Country, error) {
	countries := make([]core.Country, 0, len(storedCountries))

	for _, country := range storedCountries {
		countries = append(countries, country)
	}

	return countries, nil
}
