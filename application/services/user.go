package services

import (
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/infrastructure/logging"
)

// UserService implements all the usecases related to one of the application users
type UserService struct {
	users  repositories.UsersRepository
	logger logging.Logger
}

// NewUserService creates a new UserService
func NewUserService(repository repositories.UsersRepository, logger logging.Logger) (*UserService, error) {
	return &UserService{users: repository, logger: logger}, nil
}

// CreateUser creates a new entity.User with a given username and password. If the username has already been taken,
// an application.ErrUsernameAlreadyInUse will be returned instead.
func (s *UserService) CreateUser(username, password string) (*entity.User, error) {
	_, err := s.users.FindUserByUsername(username)
	if err == entity.ErrNotFound {
		user, err := entity.NewUser(username, password)
		if err != nil {
			s.logger.Debug(application.FailedToCreateUser, logging.Error(err))
			return nil, err
		}

		if err := s.users.CreateUser(user); err != nil {
			s.logger.Error(application.FailedToCreateUser, err)
			return nil, application.ErrInternalError
		}

		s.logger.Debug(application.UserCreated, logging.String("user_id", user.ID))
		return user, nil
	}

	if err != nil {
		s.logger.Error(application.FailedToCreateUser, err)
		return nil, application.ErrInternalError
	}

	s.logger.Info(application.FailedToCreateUser, logging.Error(application.ErrUsernameAlreadyInUse))
	return nil, application.ErrUsernameAlreadyInUse
}

// FindUserByID returns an entity.User with a given userID. If no entities are found, entity.ErrNotFound
// should be returned instead.
func (s *UserService) FindUserByID(userID string) (*entity.User, error) {
	return s.users.FindUserByID(userID)
}

// ListUsers returns a list with all known users
func (s *UserService) ListUsers() ([]*entity.User, error) { return s.users.ListUsers() }
