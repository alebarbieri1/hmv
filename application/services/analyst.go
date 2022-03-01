package services

import (
	"flavioltonon/hmv/application"
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/repositories"
	"flavioltonon/hmv/domain/valueobject"
	"flavioltonon/hmv/infrastructure/logging"
)

type AnalystService struct {
	analysts repositories.AnalystsRepository
	users    repositories.UsersRepository
	logger   logging.Logger
}

func NewAnalystService(repository repositories.AnalystsRepository, logger logging.Logger) (*AnalystService, error) {
	return &AnalystService{analysts: repository, logger: logger}, nil
}

func (s *AnalystService) CreateAnalyst(userID string) (*entity.Analyst, error) {
	user, err := s.users.FindUserByID(userID)
	if err != nil {
		s.logger.Info(application.FailedToFindUser, logging.Error(err))
		return nil, err
	}

	// If the user does not have the input profile yet, we save it to its data
	if err := user.AddProfile(valueobject.AnalystProfile); err == nil {
		if err := s.users.UpdateUser(user); err != nil {
			s.logger.Error(application.FailedToUpdateUser, err)
			return nil, application.ErrInternalError
		}
	}

	// No matter if the user is set with an analyst profile or not - at this point we can simply try to find its
	// data:
	// - If it is not found, we should created it;
	// - If we find it, we should return an error to the user
	_, err = s.analysts.FindAnalystByUserID(userID)
	if err == entity.ErrNotFound {
		analyst, err := entity.NewAnalyst(userID)
		if err != nil {
			s.logger.Info(application.FailedToCreateAnalyst, logging.Error(err))
			return nil, err
		}

		if err := s.analysts.CreateAnalyst(analyst); err != nil {
			s.logger.Error(application.FailedToCreateAnalyst, err)
			return nil, application.ErrInternalError
		}

		s.logger.Debug(
			application.AnalystCreated,
			logging.String("user_id", userID),
			logging.String("analyst_id", analyst.ID),
		)

		return analyst, nil
	}

	if err != nil {
		s.logger.Error(application.FailedToCreateAnalyst, err)
		return nil, application.ErrInternalError
	}

	s.logger.Info(
		application.FailedToCreateAnalyst,
		logging.String("user_id", userID),
		logging.Error(application.ErrUserAlreadyIsAnAnalyst),
	)
	return nil, application.ErrUserAlreadyIsAnAnalyst
}
