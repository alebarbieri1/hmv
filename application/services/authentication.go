package services

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
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
		return nil, entity.ErrInvalidUsernameOrPassword
	}

	if user.Password != password {
		return nil, entity.ErrInvalidUsernameOrPassword
	}

	return user, nil
}
