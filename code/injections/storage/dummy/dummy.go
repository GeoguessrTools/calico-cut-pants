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

func (d dummyImpl) Get(countryId string) (*core.Country, error) {
	if country, ok := storedCountries[countryId]; ok {
		return &country, nil
	}

	return nil, fmt.Errorf("could not find country %s", countryId)
}
