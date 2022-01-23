package usecases

import "flavioltonon/hmv/domain/entity"

type AuthenticationUsecase interface {
	AuthenticateUser(username, password string) (*entity.User, error)
}
