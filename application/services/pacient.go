package services

import (
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/domain/valueobject"
	"flavioltonon/hmv/infrastructure/errors"
	"flavioltonon/hmv/infrastructure/logging"
)

// PacientService implements all the usecases related to a Pacient
type PacientService struct {
	pacients repositories.PacientsRepository
	users    repositories.UsersRepository
	logger   logging.Logger
}

// NewPacientService creates a new PacientService
func NewPacientService(
	pacients repositories.PacientsRepository,
	users repositories.UsersRepository,
	logger logging.Logger,
) *PacientService {
	return &PacientService{pacients: pacients, users: users, logger: logger}
}

// CreatePacient creates a new entity.Pacient
func (s *PacientService) CreatePacient(userID string) (*entity.Pacient, error) {
	user, err := s.users.FindUserByID(userID)
	if err != nil {
		s.logger.Error(application.FailedToFindUser, err, logging.String("user_id", userID))
		return nil, err
	}

	if !user.IsPacient() {
		if err := user.SetProfileKind(valueobject.Pacient_ProfileKind); err != nil {
			s.logger.Info(application.FailedToUpdateUser, logging.Error(application.ErrInvalidUserProfile))
			return nil, application.ErrInvalidUserProfile
		}

		if err := s.users.UpdateUser(user); err != nil {
			s.logger.Error(application.FailedToUpdateUser, err)
			return nil, application.ErrInternalError
		}
	}

	// No matter if the user is set with an pacient profile or not - at this point we can simply try to find its
	// data:
	// - If it is not found, we should created it;
	// - If we find it, we should return an error to the user
	pacient, err := s.pacients.FindPacientByUserID(user.ID)
	if err == nil {
		s.logger.Debug(application.FailedToCreatePacient, logging.Error(application.ErrUserAlreadyIsAPacient))
		return pacient, application.ErrUserAlreadyIsAPacient
	}

	if err != nil && err != entity.ErrNotFound {
		s.logger.Error(application.FailedToCreateRescuer, err)
		return nil, application.ErrInternalError
	}

	pacient, err = user.NewPacient()
	if err != nil {
		s.logger.Debug(application.FailedToCreatePacient, logging.Error(err))
		return nil, err
	}

	if err := s.pacients.CreatePacient(pacient); err != nil {
		s.logger.Error(application.FailedToCreatePacient, err)
		return nil, application.ErrInternalError
	}

	s.logger.Debug(
		application.RescuerCreated,
		logging.String("user_id", user.ID),
		logging.String("pacient_id", pacient.ID),
	)

	return pacient, nil
}

// FindPacientByID returns an entity.Pacient with a given pacientID. If no entities are found, entity.ErrNotFound
// should be returned instead.
func (s *PacientService) FindPacientByID(pacientID string) (*entity.Pacient, error) {
	return s.pacients.FindPacientByID(pacientID)
}

// FindPacientByUserID returns an entity.Pacient with a given userID. If no entities are found, entity.ErrNotFound
// should be returned instead.
func (s *PacientService) FindPacientByUserID(userID string) (*entity.Pacient, error) {
	return s.pacients.FindPacientByUserID(userID)
}

// UpdateEmergencyContact updates the EmergencyContact of a entity.Pacient with a given pacientID. This action can only
// be performed by users with an Analyst_ProfileKind.
func (s *PacientService) UpdateEmergencyContact(userID, pacientID string, emergencyContact valueobject.EmergencyContact) (*entity.Pacient, error) {
	user, err := s.users.FindUserByID(userID)
	if err != nil {
		s.logger.Error(application.FailedToFindUser, err, logging.String("user_id", userID))
		return nil, errors.WithMessage(application.FailedToFindUser, err)
	}

	if !user.IsPacient() {
		s.logger.Info(application.FailedToUpdateEmergency, logging.Error(application.ErrInvalidUserProfile), logging.String("user_id", userID))
		return nil, application.ErrInvalidUserProfile
	}

	pacient, err := s.pacients.FindPacientByUserID(user.ID)
	if err != nil {
		s.logger.Error(application.FailedToFindPacient, err, logging.String("pacient_id", pacientID))
		return nil, errors.WithMessage(application.FailedToFindPacient, err)
	}

	if pacient.ID != pacientID {
		s.logger.Info(application.FailedToFindPacient,
			logging.Error(entity.ErrNotFound),
			logging.String("user_id", userID),
			logging.String("pacient_id", pacientID),
		)
		return nil, errors.WithMessage(application.FailedToFindPacient, entity.ErrNotFound)
	}

	if err := pacient.UpdateEmergencyContact(emergencyContact); err != nil {
		s.logger.Debug(application.FailedToUpdatePacient, logging.Error(err))
		return nil, err
	}

	if err := s.pacients.UpdatePacient(pacient); err != nil {
		s.logger.Error(application.FailedToUpdatePacient, err)
		return nil, application.ErrInternalError
	}

	s.logger.Debug(application.PacientUpdated, logging.String("pacient_id", pacient.ID))
	return pacient, nil
}
