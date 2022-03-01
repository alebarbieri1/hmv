package middleware

import (
	"net/http"

	"flavioltonon/hmv/application"
	"flavioltonon/hmv/application/usecases"
	"flavioltonon/hmv/infrastructure/context"
	"flavioltonon/hmv/infrastructure/logging"
	"flavioltonon/hmv/infrastructure/presenter"
	"flavioltonon/hmv/infrastructure/response"

	"github.com/justinas/alice"
)

func Authentication(authentication usecases.AuthenticationUsecase, logger logging.Logger, presenter presenter.Presenter) alice.Constructor {
	return alice.Constructor(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := authentication.AuthenticateUserFromRequest(r)
			if err != nil {
				logger.Info(application.FailedToAuthenticateUser, logging.Error(err))
				presenter.Present(w, response.Unauthorized(application.FailedToAuthenticateUser, err))
				return
			}

			ctx := context.Background()
			ctx.Add(context.UserKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	})
}
