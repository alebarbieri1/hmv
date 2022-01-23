package middleware

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/urfave/negroni"
)

func ResponseWrapper() alice.Constructor {
	return alice.Constructor(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			wrapper := negroni.NewResponseWriter(w)
			next.ServeHTTP(wrapper, r)
		})
	})
}
