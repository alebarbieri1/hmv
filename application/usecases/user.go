package usecases

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
)

type UserUsecase interface {
	CreateUser(username, password string) (*entity.User, error)
	FindUserByID(userID string) (*entity.User, error)
	AddProfileToUser(userID string, profile valueobject.Profile) (*entity.User, error)
}
