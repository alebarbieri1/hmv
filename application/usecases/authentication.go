package usecases

import (
	"flavioltonon/hmv/domain/entity"
	"net/http"
)

// AuthenticationService defines all the usecases related to users authentication
type AuthenticationUsecase interface {
	// AuthenticateUser validates a username/password pair and returns its related entity.User, in case it exists
	AuthenticateUser(username, password string) (*entity.User, error)

	// AuthenticateUser validates a username/password pair from a http.Request and returns its related entity.User, in case it exists
	AuthenticateUserFromRequest(r *http.Request) (*entity.User, error)
}
