package v1

import (
	"encoding/json"
	"fmt"
	"github.com/GeoguessrTools/calico-cut-pants/core"
	"github.com/gorilla/mux"
	"net/http"
)

type routerImpl struct {
	calicoCut core.CalicoCutService
}

func NewHttpRouter(calicoCut core.CalicoCutService) *mux.Router {
	router := mux.NewRouter()

	impl := routerImpl{
		calicoCut: calicoCut,
	}

	router.HandleFunc("/countries/{id}", impl.HandleSingleCountry)

	return router
}

func (ri *routerImpl) HandleSingleCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryId := vars["id"]

	country, err := ri.calicoCut.Get(countryId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("failed to retrieve country %s: %s", countryId, err.Error())))
		return
	}

	decode, err := json.Marshal(country)
	w.Write(decode)
}
