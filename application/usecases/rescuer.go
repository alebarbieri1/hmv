package usecases

import "flavioltonon/hmv/domain/entity"

type RescuerUsecase interface {
	CreateRescuer(user *entity.User) (*entity.Rescuer, error)
	FindRescuerByID(rescuerID string) (*entity.Rescuer, error)
}
