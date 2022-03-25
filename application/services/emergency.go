package services

import (
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/domain/valueobject"
	"flavioltonon/hmv/infrastructure/errors"
	"flavioltonon/hmv/infrastructure/logging"
)

// EmergencyService implements all the usecases related to pacients emergencies
type EmergencyService struct {
	emergencies repositories.EmergenciesRepository
	pacients    repositories.PacientsRepository
	users       repositories.UsersRepository
	logger      logging.Logger
}

// NewEmergencyService creates a new EmergencyService
func NewEmergencyService(
	emergencies repositories.EmergenciesRepository,
	pacients repositories.PacientsRepository,
	users repositories.UsersRepository,
	logger logging.Logger,
) (*EmergencyService, error) {
	return &EmergencyService{
		emergencies: emergencies,
		pacients:    pacients,
		users:       users,
		logger:      logger,
	}, nil
}

// CreateEmergency creates a new entity.Emergency
func (s *EmergencyService) CreateEmergency(userID string) (*entity.Emergency, error) {
	user, err := s.users.FindUserByID(userID)
	if err != nil {
		s.logger.Error(application.FailedToFindUser, err, logging.String("user_id", userID))
		return nil, err
	}

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

// UpdateEmergencyForm updates the EmergencyForm of a entity.Emergency with a given emergencyID. This action can only
// be performed by users with Pacient_ProfileKind or Rescuer_ProfileKind.
func (s *EmergencyService) UpdateEmergencyForm(userID string, emergencyID string, form valueobject.EmergencyForm) (*entity.Emergency, error) {
	user, err := s.users.FindUserByID(userID)
	if err != nil {
		s.logger.Error(application.FailedToFindUser, err, logging.String("user_id", userID))
		return nil, err
	}

	emergency, err := s.FindEmergencyByID(emergencyID)
	if err != nil {
		s.logger.Error(application.FailedToFindEmergency, err)
		return nil, err
	}

	switch {
	case user.IsPacient():
		pacient, err := s.pacients.FindPacientByUserID(userID)
		if err != nil {
			return nil, err
		}

		if pacient.ID != emergency.PacientID {
			s.logger.Info(application.FailedToFindEmergency,
				logging.Error(entity.ErrNotFound),
				logging.String("user_id", userID),
				logging.String("emergency_id", emergency.ID),
			)
			return nil, errors.WithMessage(application.FailedToFindEmergency, entity.ErrNotFound)

		}
	case user.IsRescuer():
		break
	default:
		s.logger.Debug(application.FailedToListEmergencies, logging.Error(application.ErrUserMustBeAPacient))
		return nil, application.ErrUserMustBeAPacientOrRescuer
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

// FindEmergencyByID returns an entity.Emergency with a given emergencyID. If no entities are found, entity.ErrNotFound
// should be returned instead.
func (s *EmergencyService) FindEmergencyByID(emergencyID string) (*entity.Emergency, error) {
	return s.emergencies.FindEmergencyByID(emergencyID)
}

// ListEmergencies returns a list with all known entity.Emergency entities
func (s *EmergencyService) ListEmergencies() ([]*entity.Emergency, error) {
	return s.emergencies.ListEmergencies()
}

// ListEmergenciesByStatus returns a list with all known entity.Emergency entities that have a given valueobject.EmergencyStatus
func (s *EmergencyService) ListEmergenciesByStatus(status valueobject.EmergencyStatus) ([]*entity.Emergency, error) {
	return s.emergencies.ListEmergenciesByStatus(status)
}

// ListEmergenciesByStatus returns a list with all known entity.Emergency entities that are related to a User
func (s *EmergencyService) ListEmergenciesByUser(userID string) ([]*entity.Emergency, error) {
	user, err := s.users.FindUserByID(userID)
	if err != nil {
		s.logger.Error(application.FailedToFindUser, err, logging.String("user_id", userID))
		return nil, err
	}

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

// UpdateEmergencyStatus updates the EmergencyStatus of a entity.Emergency with a given emergencyID.
func (s *EmergencyService) UpdateEmergencyStatus(emergencyID string, status valueobject.EmergencyStatus) (*entity.Emergency, error) {
	emergency, err := s.emergencies.FindEmergencyByID(emergencyID)
	if err != nil {
		return nil, err
	}

	if err := emergency.UpdateStatus(status); err != nil {
		s.logger.Debug(application.FailedToUpdateEmergency, logging.Error(err))
		return nil, err
	}

	if err := s.emergencies.UpdateEmergency(emergency); err != nil {
		s.logger.Error(application.FailedToUpdateEmergency, err)
		return nil, application.ErrInternalError
	}

	s.logger.Debug(application.EmergencyUpdated, logging.String("emergency_id", emergency.ID))
	return emergency, nil
}

// SendAmbulance updates the EmergencyStatus of a entity.Emergency with a given emergencyID to AmbulanceToPacient_EmergencyStatus.
func (s *EmergencyService) SendAmbulance(userID string, emergencyID string) (*entity.Emergency, error) {
	user, err := s.users.FindUserByID(userID)
	if err != nil {
		s.logger.Error(application.FailedToFindUser, err, logging.String("user_id", userID))
		return nil, err
	}

	if !user.IsAnalyst() {
		s.logger.Debug(application.FailedToSendAmbulance, logging.Error(application.ErrUserMustBeAnAnalyst))
		return nil, application.ErrUserMustBeAnAnalyst
	}

	return s.UpdateEmergencyStatus(emergencyID, valueobject.AmbulanceToPacient_EmergencyStatus)
}

// RemovePacient updates the EmergencyStatus of a entity.Emergency with a given emergencyID to AmbulanceToHospital_EmergencyStatus.
func (s *EmergencyService) RemovePacient(userID string, emergencyID string) (*entity.Emergency, error) {
	user, err := s.users.FindUserByID(userID)
	if err != nil {
		s.logger.Error(application.FailedToFindUser, err, logging.String("user_id", userID))
		return nil, err
	}

	if !user.IsRescuer() {
		s.logger.Debug(application.FailedToRemovePacient, logging.Error(application.ErrUserMustBeARescuer))
		return nil, application.ErrUserMustBeARescuer
	}

	return s.UpdateEmergencyStatus(emergencyID, valueobject.AmbulanceToHospital_EmergencyStatus)
}

// FinishEmergencyCare updates the EmergencyStatus of a entity.Emergency with a given emergencyID to Finished_EmergencyStatus.
func (s *EmergencyService) FinishEmergencyCare(userID string, emergencyID string) (*entity.Emergency, error) {
	user, err := s.users.FindUserByID(userID)
	if err != nil {
		s.logger.Error(application.FailedToFindUser, err, logging.String("user_id", userID))
		return nil, err
	}

	if !user.IsAnalyst() {
		s.logger.Debug(application.FailedToFinishEmergencyCare, logging.Error(application.ErrUserMustBeAnAnalyst))
		return nil, application.ErrUserMustBeAnAnalyst
	}

	return s.UpdateEmergencyStatus(emergencyID, valueobject.Finished_EmergencyStatus)
}

// CancelEmergency updates the EmergencyStatus of a entity.Emergency with a given emergencyID to Cancelled_EmergencyStatus.
func (s *EmergencyService) CancelEmergency(emergencyID string) (*entity.Emergency, error) {
	return s.UpdateEmergencyStatus(emergencyID, valueobject.Cancelled_EmergencyStatus)
}
