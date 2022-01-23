package services

import (
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
)

type EmergencyService struct {
	emergencies repositories.EmergenciesRepository
	pacients    repositories.PacientsRepository
}

func NewEmergencyService(
	emergencies repositories.EmergenciesRepository,
	pacients repositories.PacientsRepository,
) (*EmergencyService, error) {
	return &EmergencyService{
		emergencies: emergencies,
		pacients:    pacients,
	}, nil
}

func (s *EmergencyService) CreateEmergency(userID string) (*entity.Emergency, error) {
	pacient, err := s.pacients.FindPacientByUserID(userID)
	if err == entity.ErrNotFound {
		return nil, application.ErrUserMustBeAPacient
	}

	if err != nil {
		return nil, err
	}

	emergency, err := pacient.CreateEmergency()
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
