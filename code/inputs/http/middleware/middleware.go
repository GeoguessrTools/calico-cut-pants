package middleware

import (
	"net/http"
)

// Middleware defines a middleware function.
type Middleware func(http.HandlerFunc) http.HandlerFunc

// BuildChain builds the middleware chain.
func BuildChain(f http.HandlerFunc, m ...Middleware) http.HandlerFunc {
	if len(m) == 0 {
		return f
	}
	return m[0](BuildChain(f, m[1:cap(m)]...))
}
