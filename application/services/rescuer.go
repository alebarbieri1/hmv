package services

import (
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/domain/valueobject"
	"flavioltonon/hmv/infrastructure/logging"
)

type RescuerService struct {
	rescuers repositories.RescuersRepository
	users    repositories.UsersRepository
	logger   logging.Logger
}

func NewRescuerService(
	rescuers repositories.RescuersRepository,
	users repositories.UsersRepository,
	logger logging.Logger,
) (*RescuerService, error) {
	return &RescuerService{rescuers: rescuers, users: users, logger: logger}, nil
}

func (s *RescuerService) CreateRescuer(user *entity.User) (*entity.Rescuer, error) {
	// If the user does not have the input profile yet, we save it to its data
	if err := user.SetProfileKind(valueobject.Rescuer_ProfileKind); err == nil {
		if err := s.users.UpdateUser(user); err != nil {
			s.logger.Error(application.FailedToUpdateUser, err)
			return nil, application.ErrInternalError
		}
	} else {
		s.logger.Debug(application.FailedToUpdateUser, logging.Error(err))
	}

	// No matter if the user is set with an rescuer profile or not - at this point we can simply try to find its
	// data:
	// - If it is not found, we should created it;
	// - If we find it, we should return an error to the user
	_, err := s.rescuers.FindRescuerByUserID(user.ID)
	if err == entity.ErrNotFound {
		rescuer, err := entity.NewRescuer(user.ID)
		if err != nil {
			s.logger.Debug(application.FailedToCreateRescuer, logging.Error(err))
			return nil, err
		}

		if err := s.rescuers.CreateRescuer(rescuer); err != nil {
			s.logger.Error(application.FailedToCreateRescuer, err)
			return nil, application.ErrInternalError
		}

		s.logger.Debug(
			application.RescuerCreated,
			logging.String("user_id", user.ID),
			logging.String("rescuer_id", rescuer.ID),
		)

		return rescuer, nil
	}

	if err != nil {
		s.logger.Error(application.FailedToCreateRescuer, err)
		return nil, application.ErrInternalError
	}

	return nil, application.ErrUserAlreadyIsARescuer
}

func (s *RescuerService) FindRescuerByID(rescuerID string) (*entity.Rescuer, error) {
	return s.rescuers.FindRescuerByID(rescuerID)
}
