package v

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/GeoguessrTools/calico-cut-pants/core"
	"github.com/GeoguessrTools/calico-cut-pants/inputs/http/errormapper"
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
		errCode := errormapper.GetResponseCodeFromError(err)
		w.WriteHeader(errCode)
		w.Write([]byte(err.Error()))
		return
	}

	decode, err := json.Marshal(country)
	if err != nil {
		// TODO: build out a more robust error reporting for common http issues.
		//       this is something we will need, especially once we start handling user input.
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("failed to generate response for country %s: %s", countryId, err.Error())))
		return
	}
	w.Write(decode)
}
