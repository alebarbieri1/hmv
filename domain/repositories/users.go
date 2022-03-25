package repositories

import "flavioltonon/hmv/domain/entity"

type UsersRepository interface {
	CreateUser(user *entity.User) error
	FindUserByID(userID string) (*entity.User, error)
	FindUserByUsername(username string) (*entity.User, error)

	// ListUsers returns all the entity.User in the repository
	ListUsers() ([]*entity.User, error)

	UpdateUser(user *entity.User) error
}
