package usecases

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
)

type EmergencyUsecase interface {
	CreateEmergency(user *entity.User) (*entity.Emergency, error)
	FindEmergencyByID(emergencyID string) (*entity.Emergency, error)
	ListEmergencies() ([]*entity.Emergency, error)
	ListEmergenciesByStatus(status valueobject.EmergencyStatus) ([]*entity.Emergency, error)
	ListUserEmergencies(user *entity.User) ([]*entity.Emergency, error)
	UpdateEmergencyStatus(emergency *entity.Emergency, status valueobject.EmergencyStatus) error
	StartEmergencyCare(user *entity.User, emergency *entity.Emergency) error
	CancelEmergency(emergency *entity.Emergency) error
}
