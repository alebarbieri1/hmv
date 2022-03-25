package usecases

import (
	"flavioltonon/hmv/domain/entity"
)

// UserService defines all the usecases related to one of the application users
type UserUsecase interface {
	// CreateUser creates a new entity.User with a given username and password. If the username has already been taken,
	// an application.ErrUsernameAlreadyInUse will be returned instead.
	CreateUser(username, password string) (*entity.User, error)

	// FindUserByID returns an entity.User with a given userID. If no entities are found, entity.ErrNotFound
	// should be returned instead.
	FindUserByID(userID string) (*entity.User, error)

	// ListUsers returns a list with all known users
	ListUsers() ([]*entity.User, error)
}
