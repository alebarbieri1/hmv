package usecases

import "flavioltonon/hmv/domain/entity"

type EmergencyUsecase interface {
	CreateEmergency(userID string) (*entity.Emergency, error)
	ListEmergencies() ([]*entity.Emergency, error)
}
