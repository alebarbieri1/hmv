package services

import (
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/infrastructure/logging"
	"net/http"
)

type AuthenticationService struct {
	users  repositories.UsersRepository
	logger logging.Logger
}

func NewAuthenticationService(repository repositories.UsersRepository, logger logging.Logger) (*AuthenticationService, error) {
	return &AuthenticationService{users: repository, logger: logger}, nil
}

func (s *AuthenticationService) AuthenticateUser(username, password string) (*entity.User, error) {
	user, err := s.users.FindUserByUsername(username)
	if err != nil {
		s.logger.Info(application.FailedToAuthenticateUser, logging.Error(application.ErrInvalidUsername))
		return nil, application.ErrInvalidUsernameOrPassword
	}

	if user.Password != password {
		s.logger.Info(application.FailedToAuthenticateUser, logging.Error(application.ErrInvalidPassword))
		return nil, application.ErrInvalidUsernameOrPassword
	}

	s.logger.Debug(application.UserAuthenticated, logging.String("username", username))
	return user, nil
}

func (s *AuthenticationService) AuthenticateUserFromRequest(r *http.Request) (*entity.User, error) {
	username, password, hasCredentials := r.BasicAuth()
	if !hasCredentials {
		s.logger.Info(application.FailedToAuthenticateUser, logging.Error(application.ErrBasicAuthenticationRequired))
		return nil, application.ErrBasicAuthenticationRequired
	}

	return s.AuthenticateUser(username, password)
}
