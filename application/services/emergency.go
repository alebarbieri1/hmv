package services

import (
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/infrastructure/logging"
)

type EmergencyService struct {
	emergencies repositories.EmergenciesRepository
	pacients    repositories.PacientsRepository
	logger      logging.Logger
}

func NewEmergencyService(
	emergencies repositories.EmergenciesRepository,
	pacients repositories.PacientsRepository,
	logger logging.Logger,
) (*EmergencyService, error) {
	return &EmergencyService{
		emergencies: emergencies,
		pacients:    pacients,
		logger:      logger,
	}, nil
}

func (s *EmergencyService) CreateEmergency(userID string) (*entity.Emergency, error) {
	pacient, err := s.pacients.FindPacientByUserID(userID)
	if err == entity.ErrNotFound {
		s.logger.Info(
			application.FailedToCreateEmergency,
			logging.Error(application.ErrUserMustBeAPacient),
			logging.String("user_id", userID),
		)
		return nil, application.ErrUserMustBeAPacient
	}

	if err != nil {
		s.logger.Error(application.FailedToCreateEmergency, err)
		return nil, application.ErrInternalError
	}

	emergency, err := entity.NewEmergency(pacient.ID)
	if err != nil {
		s.logger.Info(application.FailedToCreateEmergency, logging.Error(err))
		return nil, err
	}

	if err := s.emergencies.CreateEmergency(emergency); err != nil {
		s.logger.Error(application.FailedToCreateEmergency, err)
		return nil, application.ErrInternalError
	}

	s.logger.Debug(
		application.EmergencyCreated,
		logging.String("emergency_id", emergency.ID),
		logging.String("user_id", userID),
	)
	return emergency, nil
}

func (s *EmergencyService) ListEmergencies() ([]*entity.Emergency, error) {
	return s.emergencies.ListEmergencies()
}

func (s *EmergencyService) ListEmergenciesByPacientID(pacientID string) ([]*entity.Emergency, error) {
	return s.emergencies.ListEmergenciesByPacientID(pacientID)
}
