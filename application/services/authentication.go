package services

import (
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"net/http"
)

type AuthenticationService struct {
	users repositories.UsersRepository
}

func NewAuthenticationService(repository repositories.UsersRepository) (*AuthenticationService, error) {
	return &AuthenticationService{users: repository}, nil
}

func (s *AuthenticationService) AuthenticateUser(username, password string) (*entity.User, error) {
	user, err := s.users.FindUserByUsername(username)
	if err != nil {
		return nil, application.ErrInvalidUsernameOrPassword
	}

	if user.Password != password {
		return nil, application.ErrInvalidUsernameOrPassword
	}

	return user, nil
}

func (s *AuthenticationService) AuthenticateUserFromRequest(r *http.Request) (*entity.User, error) {
	username, password, hasCredentials := r.BasicAuth()
	if !hasCredentials {
		return nil, application.ErrBasicAuthenticationRequired
	}

	return s.AuthenticateUser(username, password)
}
