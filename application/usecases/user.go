package usecases

import "flavioltonon/hmv/domain/entity"

type UserUsecase interface {
	CreateUser(username, password string) (*entity.User, error)
	FindUser(userID string) (*entity.User, error)
}
