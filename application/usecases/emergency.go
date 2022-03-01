package usecases

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
)

type EmergencyUsecase interface {
	CreateEmergency(user *entity.User) (*entity.Emergency, error)
	ListEmergencies() ([]*entity.Emergency, error)
	ListEmergenciesByStatus(status valueobject.EmergencyStatus) ([]*entity.Emergency, error)
	ListEmergenciesByPacient(pacient *entity.Pacient) ([]*entity.Emergency, error)
}
