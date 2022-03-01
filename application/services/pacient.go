package services

import (
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/domain/valueobject"
	"flavioltonon/hmv/infrastructure/logging"
)

type PacientService struct {
	pacients repositories.PacientsRepository
	users    repositories.UsersRepository
	logger   logging.Logger
}

func NewPacientService(
	pacients repositories.PacientsRepository,
	users repositories.UsersRepository,
	logger logging.Logger,
) (*PacientService, error) {
	return &PacientService{pacients: pacients, users: users, logger: logger}, nil
}

func (s *PacientService) CreatePacient(user *entity.User) (*entity.Pacient, error) {
	// If the user does not have the input profile yet, we save it to its data
	if err := user.SetProfileKind(valueobject.Pacient_ProfileKind); err == nil {
		s.logger.Debug(
			application.UserCanBeUpdated,
			logging.String("user_id", user.ID),
			logging.Stringer("profile", valueobject.Pacient_ProfileKind),
		)

		if err := s.users.UpdateUser(user); err != nil {
			s.logger.Error(application.FailedToUpdateUser, err)
			return nil, application.ErrInternalError
		}

		s.logger.Debug(
			application.UserUpdated,
			logging.String("user_id", user.ID),
			logging.Stringer("profile", valueobject.Pacient_ProfileKind),
		)
	}

	// No matter if the user is set with an pacient profile or not - at this point we can simply try to find its
	// data:
	// - If it is not found, we should created it;
	// - If we find it, we should return an error to the user
	_, err := s.pacients.FindPacientByUserID(user.ID)
	if err == entity.ErrNotFound {
		pacient, err := entity.NewPacient(user.ID)
		if err != nil {
			s.logger.Debug(application.FailedToCreatePacient, logging.Error(err))
			return nil, err
		}

		if err := s.pacients.CreatePacient(pacient); err != nil {
			s.logger.Error(application.FailedToCreatePacient, err)
			return nil, application.ErrInternalError
		}

		s.logger.Debug(
			application.PacientCreated,
			logging.String("user_id", user.ID),
			logging.String("pacient_id", pacient.ID),
		)

		return pacient, nil
	}

	if err != nil {
		s.logger.Error(application.FailedToCreatePacient, err)
		return nil, application.ErrInternalError
	}

	return nil, application.ErrUserAlreadyIsAPacient
}

func (s *PacientService) FindPacientByID(pacientID string) (*entity.Pacient, error) {
	return s.pacients.FindPacientByID(pacientID)
}

func (s *PacientService) FindPacientByUserID(userID string) (*entity.Pacient, error) {
	return s.pacients.FindPacientByUserID(userID)
}

func (s *PacientService) UpdateEmergencyContact(pacient *entity.Pacient, emergencyContact valueobject.EmergencyContact) (*entity.Pacient, error) {
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
