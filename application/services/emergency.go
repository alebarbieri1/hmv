package services

import (
	"context"
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

func (s *EmergencyService) UpdateEmergencyForm(ctx context.Context, user *entity.User, emergencyID string, form valueobject.EmergencyForm) (*entity.Emergency, error) {
	if !user.IsPacient() && !user.IsRescuer() {
		s.logger.Debug(application.FailedToListEmergencies, logging.Error(application.ErrUserMustBeAPacient))
		return nil, application.ErrUserMustBeAPacientOrRescuer
	}

	emergency, err := s.FindEmergencyByID(emergencyID)
	if err != nil {
		s.logger.Error(application.FailedToFindEmergency, err)
		return nil, err
	}

	if err := emergency.UpdateForm(form); err != nil {
		s.logger.Error(application.FailedToUpdateEmergencyForm, err)
		return nil, err
	}

	if err := s.emergencies.UpdateEmergency(emergency); err != nil {
		s.logger.Error(application.FailedToUpdateEmergency, err)
		return nil, application.ErrInternalError
	}

	return emergency, nil
}

func (s *EmergencyService) FindEmergencyByID(emergencyID string) (*entity.Emergency, error) {
	return s.emergencies.FindEmergencyByID(emergencyID)
}

func (s *EmergencyService) ListEmergencies() ([]*entity.Emergency, error) {
	return s.emergencies.ListEmergencies()
}

func (s *EmergencyService) ListEmergenciesByStatus(status valueobject.EmergencyStatus) ([]*entity.Emergency, error) {
	return s.emergencies.ListEmergenciesByStatus(status)
}

func (s *EmergencyService) ListUserEmergencies(user *entity.User) ([]*entity.Emergency, error) {
	if !user.IsPacient() {
		s.logger.Debug(application.FailedToListEmergencies, logging.Error(application.ErrUserMustBeAPacient))
		return nil, application.ErrUserMustBeAPacient
	}

	pacient, err := s.pacients.FindPacientByUserID(user.ID)
	if err == entity.ErrNotFound {
		s.logger.Debug(application.FailedToFindPacient, logging.Error(err))
		return nil, application.ErrUserMustBeAPacient
	}

	if err != nil {
		s.logger.Error(application.FailedToFindPacient, err)
		return nil, application.ErrInternalError
	}

	return s.emergencies.ListEmergenciesByPacientID(pacient.ID)
}

func (s *EmergencyService) UpdateEmergencyStatus(emergency *entity.Emergency, status valueobject.EmergencyStatus) error {
	if err := emergency.UpdateStatus(status); err != nil {
		s.logger.Debug(application.FailedToUpdateEmergency, logging.Error(err))
		return err
	}

	if err := s.emergencies.UpdateEmergency(emergency); err != nil {
		s.logger.Error(application.FailedToUpdateEmergency, err)
		return application.ErrInternalError
	}

	s.logger.Debug(application.EmergencyUpdated, logging.String("emergency_id", emergency.ID))
	return nil
}

func (s *EmergencyService) SendAmbulance(user *entity.User, emergency *entity.Emergency) error {
	if !user.IsAnalyst() {
		s.logger.Debug(application.FailedToSendAmbulance, logging.Error(application.ErrUserMustBeAnAnalyst))
		return application.ErrUserMustBeAnAnalyst
	}

	return s.UpdateEmergencyStatus(emergency, valueobject.AmbulanceToPacient_EmergencyStatus)
}

func (s *EmergencyService) RemovePacient(user *entity.User, emergency *entity.Emergency) error {
	if !user.IsRescuer() {
		s.logger.Debug(application.FailedToFinishEmergencyCare, logging.Error(application.ErrUserMustBeAnAnalyst))
		return application.ErrUserMustBeAnAnalyst
	}

	return s.UpdateEmergencyStatus(emergency, valueobject.Finished_EmergencyStatus)
}

func (s *EmergencyService) FinishEmergencyCare(user *entity.User, emergency *entity.Emergency) error {
	if !user.IsAnalyst() {
		s.logger.Debug(application.FailedToFinishEmergencyCare, logging.Error(application.ErrUserMustBeAnAnalyst))
		return application.ErrUserMustBeAnAnalyst
	}

	return s.UpdateEmergencyStatus(emergency, valueobject.Finished_EmergencyStatus)
}

func (s *EmergencyService) CancelEmergency(emergency *entity.Emergency) error {
	return s.UpdateEmergencyStatus(emergency, valueobject.Cancelled_EmergencyStatus)
}
