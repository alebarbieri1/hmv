package usecases

import (
	"flavioltonon/hmv/domain/entity"
)

type UserUsecase interface {
	CreateUser(username, password string) (*entity.User, error)
	FindUserByID(userID string) (*entity.User, error)

	// ListUsers returns a list with all known users
	ListUsers() ([]*entity.User, error)
}
