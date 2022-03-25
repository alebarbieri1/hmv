package repositories

import "flavioltonon/hmv/domain/entity"

// UsersRepository is a repository for entity.User entities
type UsersRepository interface {
	// CreateUser stores an entity.User into the repository. If an User with the same ID already
	// exists in the repository, entity.ErrDuplicatedEntry should be returned instead.
	CreateUser(user *entity.User) error

	// FindUserByID returns an entity.User identified by a given userID. If no entities are found,
	// entity.ErrNotFound should be returned instead.
	FindUserByID(userID string) (*entity.User, error)

	// FindUserByUsername returns an entity.User identified by a given username. If no entities are found,
	// entity.ErrNotFound should be returned instead.
	FindUserByUsername(username string) (*entity.User, error)

	// ListUsers returns all the entity.User in the repository
	ListUsers() ([]*entity.User, error)

	// UpdateUser updates an User in the repository. If no entities with the same ID as the input entity.User are found,
	// entity.ErrNotFound should be returned instead.
	UpdateUser(user *entity.User) error
}
