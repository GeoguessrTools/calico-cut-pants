package v1

import "net/http"

type Route struct {
	Path     string
	Method   string
	Function func(w http.ResponseWriter, r *http.Request)
	// EnabledMiddleware determines which middleware should be loaded for the route. SupportCors will be loaded by
	// default.
	EnabledMiddleware []Middleware
}

func GetRoutes(hi handlerImpl) []Route {
	return []Route{
		{
			Path:              "countries/{id}",
			Method:            http.MethodGet,
			Function:          hi.HandleSingleCountry,
			EnabledMiddleware: []Middleware{},
		},
	}
}
