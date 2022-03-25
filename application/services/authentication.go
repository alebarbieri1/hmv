package services

import (
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/infrastructure/logging"
	"net/http"
)

// AuthenticationService implements all the usecases related to users authentication
type AuthenticationService struct {
	users  repositories.UsersRepository
	logger logging.Logger
}

// NewAuthenticationService creates a new AuthenticationService
func NewAuthenticationService(repository repositories.UsersRepository, logger logging.Logger) (*AuthenticationService, error) {
	return &AuthenticationService{users: repository, logger: logger}, nil
}

// AuthenticateUser validates a username/password pair and returns its related entity.User, in case it exists
func (s *AuthenticationService) AuthenticateUser(username, password string) (*entity.User, error) {
	user, err := s.users.FindUserByUsername(username)
	if err == entity.ErrNotFound {
		s.logger.Debug(application.FailedToAuthenticateUser, logging.Error(application.ErrInvalidUsername))
		return nil, application.ErrInvalidUsernameOrPassword
	}

	if err != nil {
		s.logger.Error(application.FailedToAuthenticateUser, err)
		return nil, application.ErrInternalError
	}

	if user.Password != password {
		s.logger.Debug(application.FailedToAuthenticateUser, logging.Error(application.ErrInvalidPassword))
		return nil, application.ErrInvalidUsernameOrPassword
	}

	s.logger.Debug(application.UserAuthenticated, logging.String("username", username))
	return user, nil
}

// AuthenticateUser validates a username/password pair from a http.Request and returns its related entity.User, in case it exists
func (s *AuthenticationService) AuthenticateUserFromRequest(r *http.Request) (*entity.User, error) {
	username, password, hasCredentials := r.BasicAuth()
	if !hasCredentials {
		s.logger.Debug(application.FailedToAuthenticateUser, logging.Error(application.ErrBasicAuthenticationRequired))
		return nil, application.ErrBasicAuthenticationRequired
	}

	return s.AuthenticateUser(username, password)
}
