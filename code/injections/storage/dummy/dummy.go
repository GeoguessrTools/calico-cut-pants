package dummy

import (
	"github.com/GeoguessrTools/calico-cut-pants/core"
)

type dummyImpl struct {
}

func NewDummyStorage() (core.Storage, error) {
	return &dummyImpl{}, nil
}

func (d dummyImpl) Get(countryId string) (*core.Country, error) {
	if country, ok := storedCountries[countryId]; ok {
		return &country, nil
	}

	return nil, core.Error{Code: core.CountryNotFoundErr}
}
