package middleware

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/justinas/alice"
)

func RequestID() alice.Constructor {
	return alice.Constructor(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.NewString()
			r.Header.Set("Internal-Request-Id", requestID)
			w.Header().Set("Internal-Request-Id", requestID)
			next.ServeHTTP(w, r)
		})
	})
}
