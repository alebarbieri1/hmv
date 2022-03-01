package services

import (
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/infrastructure/logging"
)

type UserService struct {
	users  repositories.UsersRepository
	logger logging.Logger
}

func NewUserService(repository repositories.UsersRepository, logger logging.Logger) (*UserService, error) {
	return &UserService{users: repository, logger: logger}, nil
}

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

func (s *UserService) FindUserByID(userID string) (*entity.User, error) {
	return s.users.FindUserByID(userID)
}
