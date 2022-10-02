package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GeoguessrTools/calico-cut-pants/core"
	"github.com/gorilla/mux"
)

type handlerImpl struct {
	calicoCut core.CalicoCutService
}

func NewHttpHandler(calicoCut core.CalicoCutService) http.Handler {
	impl := handlerImpl{
		calicoCut: calicoCut,
	}

	router := NewHttpRouter(impl)

	return router
}

func (hi *handlerImpl) HandleSingleCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryId := vars["id"]

	country, err := hi.calicoCut.GetCountry(countryId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("failed to retrieve country %s: %s", countryId, err.Error())))
		return
	}

	decode, err := json.Marshal(country)
	w.Write(decode)
}

func (hi *handlerImpl) HandleCountries(w http.ResponseWriter, r *http.Request) {
	countryNames, err := hi.calicoCut.GetAllCountryNames()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to retrieve country list: %s", err.Error())))
		return
	}

	// TODO: this should be more structured possibly.
	rv := struct {
		Countries []string `json:"countries"`
	}{
		Countries: countryNames,
	}

	decode, err := json.Marshal(rv)
	w.Write(decode)
}
