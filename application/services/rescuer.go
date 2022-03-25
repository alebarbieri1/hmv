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

func (s *RescuerService) CreateRescuer(userID string) (*entity.Rescuer, error) {
	user, err := s.users.FindUserByID(userID)
	if err != nil {
		s.logger.Error(application.FailedToFindUser, err, logging.String("user_id", userID))
		return nil, err
	}

	if !user.IsRescuer() {
		if err := user.SetProfileKind(valueobject.Rescuer_ProfileKind); err != nil {
			s.logger.Info(application.FailedToUpdateUser, logging.Error(application.ErrInvalidUserProfile))
			return nil, application.ErrInvalidUserProfile
		}

		if err := s.users.UpdateUser(user); err != nil {
			s.logger.Error(application.FailedToUpdateUser, err)
			return nil, application.ErrInternalError
		}
	}

	// No matter if the user is set with an rescuer profile or not - at this point we can simply try to find its
	// data:
	// - If it is not found, we should created it;
	// - If we find it, we should return an error to the user
	rescuer, err := s.rescuers.FindRescuerByUserID(user.ID)
	if err == nil {
		s.logger.Debug(application.FailedToCreateRescuer, logging.Error(application.ErrUserAlreadyIsARescuer))
		return rescuer, application.ErrUserAlreadyIsARescuer
	}

	if err != nil && err != entity.ErrNotFound {
		s.logger.Error(application.FailedToCreateRescuer, err)
		return nil, application.ErrInternalError
	}

	rescuer, err = entity.NewRescuer(user.ID)
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

// FindRescuerByID returns an entity.Rescuer with a given rescuerID. If no entities are found, entity.ErrNotFound
// should be returned instead.
func (s *RescuerService) FindRescuerByID(rescuerID string) (*entity.Rescuer, error) {
	return s.rescuers.FindRescuerByID(rescuerID)
}
