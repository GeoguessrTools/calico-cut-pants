package main

import (
	"fmt"
	"net/http"

	"github.com/GeoguessrTools/calico-cut-pants/core"
	"github.com/GeoguessrTools/calico-cut-pants/injections/storage/dummy"
	httpv1 "github.com/GeoguessrTools/calico-cut-pants/inputs/http/v1"
)

func main() {

	dummyStorage, _ := dummy.NewDummyStorage()

	coreService := core.NewCalicoCutService(dummyStorage)

	httpHandler := httpv1.NewHttpHandler(coreService)

	err := http.ListenAndServe(":6969", httpHandler)
	if err != nil {
		fmt.Printf("http server failed: %s", err.Error())
	}
}
