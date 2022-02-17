package middleware

import (
	"flavioltonon/hmv/infrastructure/headers"
	"flavioltonon/hmv/infrastructure/logging"
	"net/http"

	"github.com/justinas/alice"
	"github.com/urfave/negroni"
)

func Logging(logger logging.Logger) alice.Constructor {
	return alice.Constructor(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			logger.Info("new request received",
				logging.String("request_id", r.Header.Get(headers.InternalRequestID)),
				logging.String("request_method", r.Method),
				logging.String("request_uri", r.RequestURI),
				logging.String("request_source_ip", r.RemoteAddr),
			)

			next.ServeHTTP(w, r)

			if wrapper, implements := w.(negroni.ResponseWriter); implements {
				logger.Info("returning response",
					logging.String("request_id", w.Header().Get(headers.InternalRequestID)),
					logging.Int("response_status", wrapper.Status()),
				)
			}
		})
	})
}
