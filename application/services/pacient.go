package services

import (
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/domain/valueobject"
	"flavioltonon/hmv/infrastructure/errors"
	"flavioltonon/hmv/infrastructure/logging"
)

type PacientService struct {
	pacients repositories.PacientsRepository
	logger   logging.Logger
}

func NewPacientService(repository repositories.PacientsRepository, logger logging.Logger) (*PacientService, error) {
	return &PacientService{pacients: repository, logger: logger}, nil
}

func (s *PacientService) CreatePacient(userID string) (*entity.Pacient, error) {
	_, err := s.pacients.FindPacientByUserID(userID)
	if err == entity.ErrNotFound {
		pacient, err := entity.NewPacient(userID)
		if err != nil {
			s.logger.Info(application.FailedToCreatePacient, logging.Error(err))
			return nil, err
		}

		if err := s.pacients.CreatePacient(pacient); err != nil {
			s.logger.Error(application.FailedToCreatePacient, err)
			return nil, application.ErrInternalError
		}

		s.logger.Debug(
			application.PacientCreated,
			logging.String("user_id", userID),
			logging.String("pacient_id", pacient.ID),
		)
		return pacient, nil
	}

	if err != nil {
		s.logger.Error(application.FailedToCreatePacient, err)
		return nil, application.ErrInternalError
	}

	s.logger.Info(
		application.FailedToCreatePacient,
		logging.String("user_id", userID),
		logging.Error(application.ErrUserAlreadyIsAPacient),
	)
	return nil, application.ErrUserAlreadyIsAPacient
}

func (s *PacientService) FindPacientByUserID(userID string) (*entity.Pacient, error) {
	return s.pacients.FindPacientByUserID(userID)
}

func (s *PacientService) UpdateEmergencyContact(userID string, emergencyContact valueobject.EmergencyContact) (*entity.Pacient, error) {
	pacient, err := s.pacients.FindPacientByUserID(userID)
	if err != nil {
		if errors.Is(err, entity.ErrNotFound) {
			s.logger.Info(
				application.FailedToFindPacient,
				logging.Error(err),
				logging.String("user_id", userID),
			)
		} else {
			s.logger.Error(application.FailedToFindPacient, err)
		}

		return nil, errors.WithMessage(application.FailedToFindPacient, err)
	}

	if err := pacient.UpdateEmergencyContact(emergencyContact); err != nil {
		s.logger.Info(application.FailedToUpdatePacient, logging.Error(err))
		return nil, err
	}

	if err := s.pacients.UpdatePacient(pacient); err != nil {
		s.logger.Error(application.FailedToUpdatePacient, err)
		return nil, application.ErrInternalError
	}

	s.logger.Debug(application.PacientUpdated, logging.String("pacient_id", pacient.ID))
	return pacient, nil
}
