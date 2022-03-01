package repositories

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
)

type EmergenciesRepository interface {
	CreateEmergency(emergency *entity.Emergency) error
	FindEmergencyByID(emergencyID string) (*entity.Emergency, error)
	ListEmergencies() ([]*entity.Emergency, error)
	ListEmergenciesByStatus(status valueobject.EmergencyStatus) ([]*entity.Emergency, error)
	ListEmergenciesByPacientID(pacientID string) ([]*entity.Emergency, error)
	UpdateEmergency(emergency *entity.Emergency) error
}
