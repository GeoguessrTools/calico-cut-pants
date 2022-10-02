package v1

import (
	"net/http"
)

// Middleware defines middleware. It takes in a function and wraps itself around it.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// BuildChain recursively builds a middleware chain.
func BuildChain(f http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	if len(m) == 0 {
		return f
	}
	return m[0](BuildChain(f, m[1:cap(m)]...))
}

// SupportCors adds necessary headers to enable CORS on API routes.
var SupportCors = func(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		f(w, r)
	}
}
