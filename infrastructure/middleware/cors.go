package middleware

import (
	"net/http"

	"github.com/justinas/alice"
	"github.com/rs/cors"
)

// CORS is a Cross-origin resource sharing handling middleware
func CORS() alice.Constructor {
	return alice.Constructor(
		cors.New(cors.Options{
			AllowedOrigins: []string{"*"},
			AllowedMethods: []string{
				http.MethodGet,
				http.MethodPost,
				http.MethodPatch,
				http.MethodPut,
				http.MethodDelete,
				http.MethodOptions,
			},
			AllowCredentials: true,
		}).Handler,
	)
}
