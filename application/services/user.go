package services

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
)

type UserService struct {
	users repositories.UsersRepository
}

func NewUserService(repository repositories.UsersRepository) (*UserService, error) {
	return &UserService{users: repository}, nil
}

func (s *UserService) CreateUser(username, password string) (*entity.User, error) {
	_, err := s.users.FindUserByUsername(username)
	if err == entity.ErrNotFound {
		user, err := entity.NewUser(username, password)
		if err != nil {
			return nil, err
		}

		if err := s.users.CreateUser(user); err != nil {
			return nil, err
		}

		return user, nil
	}

	if err != nil {
		return nil, err
	}

	return nil, entity.ErrUsernameAlreadyInUse
}

func (s *UserService) FindUser(userID string) (*entity.User, error) {
	return s.users.FindUserByID(userID)
}
