package services

import (
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/domain/valueobject"
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

func (s *EmergencyService) CreateEmergency(user *entity.User) (*entity.Emergency, error) {
	if !user.IsPacient() {
		s.logger.Debug(
			application.FailedToCreateEmergency,
			logging.Error(application.ErrUserMustBeAPacient),
			logging.String("user_id", user.ID),
		)
		return nil, application.ErrUserMustBeAPacient
	}

	pacient, err := s.pacients.FindPacientByUserID(user.ID)
	if err == entity.ErrNotFound {
		s.logger.Debug(
			application.FailedToCreateEmergency,
			logging.Error(application.ErrUserMustBeAPacient),
			logging.String("user_id", user.ID),
		)
		return nil, application.ErrUserMustBeAPacient
	}

	if err != nil {
		s.logger.Error(application.FailedToCreateEmergency, err)
		return nil, application.ErrInternalError
	}

	emergency, err := entity.NewEmergency(pacient.ID)
	if err != nil {
		s.logger.Debug(application.FailedToCreateEmergency, logging.Error(err))
		return nil, err
	}

	if err := s.emergencies.CreateEmergency(emergency); err != nil {
		s.logger.Error(application.FailedToCreateEmergency, err)
		return nil, application.ErrInternalError
	}

	s.logger.Debug(
		application.EmergencyCreated,
		logging.String("emergency_id", emergency.ID),
		logging.String("user_id", user.ID),
	)
	return emergency, nil
}

func (s *EmergencyService) ListEmergencies() ([]*entity.Emergency, error) {
	return s.emergencies.ListEmergencies()
}

func (s *EmergencyService) ListEmergenciesByStatus(status valueobject.EmergencyStatus) ([]*entity.Emergency, error) {
	return s.emergencies.ListEmergenciesByStatus(status)
}

func (s *EmergencyService) ListEmergenciesByPacient(pacient *entity.Pacient) ([]*entity.Emergency, error) {
	return s.emergencies.ListEmergenciesByPacientID(pacient.ID)
}
