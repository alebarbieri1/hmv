package usecases

import "flavioltonon/hmv/domain/entity"

// AnalystUsecase defines all the usecases related to a hospital Analyst
type AnalystUsecase interface {
	// CreateAnalyst creates a new entity.Analyst
	CreateAnalyst(userID string) (*entity.Analyst, error)

	// FindAnalystByID returns an entity.Analyst with a given analystID. If no entities are found, entity.ErrNotFound
	// should be returned instead.
	FindAnalystByID(analystID string) (*entity.Analyst, error)
}
