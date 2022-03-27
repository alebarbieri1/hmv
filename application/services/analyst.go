package services

import (
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/domain/valueobject"
	"flavioltonon/hmv/infrastructure/logging"
)

// AnalystService implements all the usecases related to an hospital Analyst
type AnalystService struct {
	analysts repositories.AnalystsRepository
	users    repositories.UsersRepository
	logger   logging.Logger
}

// NewAnalystService creates a new AnalystService
func NewAnalystService(
	analysts repositories.AnalystsRepository,
	users repositories.UsersRepository,
	logger logging.Logger,
) *AnalystService {
	return &AnalystService{analysts: analysts, users: users, logger: logger}
}

// CreateAnalyst creates a new entity.Analyst
func (s *AnalystService) CreateAnalyst(userID string) (*entity.Analyst, error) {
	user, err := s.users.FindUserByID(userID)
	if err != nil {
		s.logger.Error(application.FailedToFindUser, err, logging.String("user_id", userID))
		return nil, err
	}

	if !user.IsAnalyst() {
		if err := user.SetProfileKind(valueobject.Analyst_ProfileKind); err != nil {
			s.logger.Info(application.FailedToUpdateUser, logging.Error(application.ErrInvalidUserProfile))
			return nil, application.ErrInvalidUserProfile
		}

		if err := s.users.UpdateUser(user); err != nil {
			s.logger.Error(application.FailedToUpdateUser, err)
			return nil, application.ErrInternalError
		}
	}

	// No matter if the user is set with an analyst profile or not - at this point we can simply try to find its
	// data:
	// - If it is not found, we should created it;
	// - If we find it, we should return an error to the user
	analyst, err := s.analysts.FindAnalystByUserID(user.ID)
	if err == nil {
		s.logger.Debug(application.FailedToCreateAnalyst, logging.Error(application.ErrUserAlreadyIsAnAnalyst))
		return analyst, application.ErrUserAlreadyIsAnAnalyst
	}

	if err != nil && err != entity.ErrNotFound {
		s.logger.Error(application.FailedToCreateAnalyst, err)
		return nil, application.ErrInternalError
	}

	analyst, err = entity.NewAnalyst(user.ID)
	if err != nil {
		s.logger.Debug(application.FailedToCreateAnalyst, logging.Error(err))
		return nil, err
	}

	if err := s.analysts.CreateAnalyst(analyst); err != nil {
		s.logger.Error(application.FailedToCreateAnalyst, err)
		return nil, application.ErrInternalError
	}

	s.logger.Debug(
		application.AnalystCreated,
		logging.String("user_id", user.ID),
		logging.String("analyst_id", analyst.ID),
	)

	return analyst, nil
}

// FindAnalystByID returns an entity.Analyst with a given analystID. If no entities are found, entity.ErrNotFound
// should be returned instead.
func (s *AnalystService) FindAnalystByID(analystID string) (*entity.Analyst, error) {
	return s.analysts.FindAnalystByID(analystID)
}
