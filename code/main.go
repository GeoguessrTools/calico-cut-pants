package main

import (
	"fmt"
	core2 "github.com/GeoguessrTools/calico-cut-pants/core"
	"github.com/GeoguessrTools/calico-cut-pants/injections/storage/dummy"
	httpv1 "github.com/GeoguessrTools/calico-cut-pants/inputs/http/v1"
	"net/http"
)

func main() {

	dummyStorage, _ := dummy.NewDummyStorage()

	core := core2.NewCalicoCutService(dummyStorage)

	httpRouter := httpv1.NewHttpRouter(core)

	err := http.ListenAndServe(":6969", httpRouter)
	if err != nil {
		fmt.Printf("http server failed: %s", err.Error())
	}
}
