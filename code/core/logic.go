package core

type CalicoCutImpl struct {
	storage Storage
}

func NewCalicoCutService(storage Storage) CalicoCutService {
	return &CalicoCutImpl{
		storage: storage,
	}
}

func (cc *CalicoCutImpl) Get(countryCode string) (*Country, error) {
	country, err := cc.storage.Get(countryCode)
	if err != nil {
		return nil, err
	}
	return country, nil
}
