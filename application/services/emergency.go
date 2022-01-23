package services

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
)

type EmergencyService struct {
	emergencies repositories.EmergenciesRepository
}

func NewEmergencyService(repository repositories.EmergenciesRepository) (*EmergencyService, error) {
	return &EmergencyService{emergencies: repository}, nil
}

func (s *EmergencyService) CreateEmergency(pacientID string) (*entity.Emergency, error) {
	emergency, err := entity.NewEmergency(pacientID)
	if err != nil {
		return nil, err
	}

	if err := s.emergencies.CreateEmergency(emergency); err != nil {
		return nil, err
	}

	return emergency, nil
}

func (s *EmergencyService) ListEmergencies() ([]*entity.Emergency, error) {
	return s.emergencies.ListEmergencies()
}

func (s *EmergencyService) ListEmergenciesByPacientID(pacientID string) ([]*entity.Emergency, error) {
	return s.emergencies.ListEmergenciesByPacientID(pacientID)
}
