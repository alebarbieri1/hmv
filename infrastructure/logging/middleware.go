package logging

import (
	"net/http"

	"github.com/justinas/alice"
)

func Middleware(logger Logger) alice.Constructor {
	return alice.Constructor(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("new request received",
				Field("request_id", r.Header.Get("Internal-Request-Id")),
				Field("request_method", r.Method),
				Field("request_uri", r.RequestURI),
				Field("request_source_ip", r.RemoteAddr),
			)

			next.ServeHTTP(w, r)

			logger.Info("returning response",
				Field("request_id", w.Header().Get("Internal-Request-Id")),
				Field("request_method", r.Method),
				Field("request_uri", r.RequestURI),
				Field("request_source_ip", r.RemoteAddr),
			)
		})
	})
}
