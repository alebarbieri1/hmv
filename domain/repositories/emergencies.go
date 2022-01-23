package repositories

import "flavioltonon/hmv/domain/entity"

type EmergenciesRepository interface {
	CreateEmergency(emergency *entity.Emergency) error
	ListEmergencies() ([]*entity.Emergency, error)
	ListEmergenciesByPacientID(pacientID string) ([]*entity.Emergency, error)
}
