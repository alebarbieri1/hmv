package middleware

import (
	"net/http"

	"flavioltonon/hmv/infrastructure/headers"

	"github.com/google/uuid"
	"github.com/justinas/alice"
)

func RequestID() alice.Constructor {
	return alice.Constructor(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := uuid.NewString()
			r.Header.Set(headers.InternalRequestID, requestID)
			w.Header().Set(headers.InternalRequestID, requestID)
			next.ServeHTTP(w, r)
		})
	})
}
