package usecases

import (
	"flavioltonon/hmv/domain/entity"
	"net/http"
)

type AuthenticationUsecase interface {
	AuthenticateUser(username, password string) (*entity.User, error)
	AuthenticateUserFromRequest(r *http.Request) (*entity.User, error)
}
