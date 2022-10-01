package v1

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GeoguessrTools/calico-cut-pants/core"
	"github.com/gorilla/mux"
)

type routerImpl struct {
	calicoCut core.CalicoCutService
}

func NewHttpRouter(calicoCut core.CalicoCutService) *mux.Router {
	router := mux.NewRouter()

	impl := routerImpl{
		calicoCut: calicoCut,
	}

	router.HandleFunc("/countries/{id}", impl.HandleSingleCountry).Methods("GET", "OPTIONS")

	return router
}

func (ri *routerImpl) HandleSingleCountry(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryId := vars["id"]

	//Allow CORS here By * or specific origin
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")

	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	country, err := ri.calicoCut.Get(countryId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("failed to retrieve country %s: %s", countryId, err.Error())))
		return
	}

	decode, err := json.Marshal(country)
	w.Write(decode)
}
