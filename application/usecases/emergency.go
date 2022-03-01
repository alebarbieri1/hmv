package usecases

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
)

type EmergencyUsecase interface {
	CreateEmergency(userID string) (*entity.Emergency, error)
	ListEmergencies() ([]*entity.Emergency, error)
	ListEmergenciesByStatus(status valueobject.EmergencyStatus) ([]*entity.Emergency, error)
	ListEmergenciesByPacientID(pacientID string) ([]*entity.Emergency, error)
}
