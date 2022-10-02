package core

type CalicoCutImpl struct {
	storage Storage
}

func NewCalicoCutService(storage Storage) CalicoCutService {
	return &CalicoCutImpl{
		storage: storage,
	}
}

func (cc *CalicoCutImpl) GetCountry(countryCode string) (*Country, error) {
	country, err := cc.storage.GetCountry(countryCode)
	if err != nil {
		return nil, err
	}
	return country, nil
}

func (cc *CalicoCutImpl) GetAllCountryNames() ([]string, error) {
	countries, err := cc.storage.GetAllCountries()
	if err != nil {
		return []string{}, err
	}

	countryNames := make([]string, 0, len(countries))
	for _, country := range countries {
		countryNames = append(countryNames, country.Name)
	}

	return countryNames, err
}
