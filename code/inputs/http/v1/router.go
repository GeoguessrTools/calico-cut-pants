package v1

import (
	"github.com/gorilla/mux"
	"net/http"
	"path"
)

const (
	apiVersionPrefix = "/v0"
)

// NewHttpRouter dynamically constructs a router from a list of Route structs. CORS headers and the OPTIONS route will
// be added to all routes.
func NewHttpRouter(hi handlerImpl) *mux.Router {
	router := mux.NewRouter()
	routes := GetRoutes(hi)

	for _, route := range routes {
		methods := []string{http.MethodOptions, route.Method}
		fullPath := path.Join(apiVersionPrefix, route.Path)
		fullMw := append(route.EnabledMiddleware, SupportCors)
		fullChain := BuildChain(route.Function, fullMw...)

		router.HandleFunc(fullPath, fullChain).Methods(methods...)
	}

	return router
}
