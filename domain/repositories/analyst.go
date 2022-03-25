package repositories

import "flavioltonon/hmv/domain/entity"

// AnalystsRepository is a repository for entity.Analyst entities
type AnalystsRepository interface {
	// CreateAnalyst stores an entity.Analyst into the repository. If an Analyst with the same ID already
	// exists in the repository, entity.ErrDuplicatedEntry should be returned instead.
	CreateAnalyst(analyst *entity.Analyst) error

	// FindAnalystByID returns an entity.Analyst identified by a given analystID. If no entities are found,
	// entity.ErrNotFound should be returned instead.
	FindAnalystByID(analystID string) (*entity.Analyst, error)

	// FindAnalystByUserID returns an entity.Analyst identified by a given userID. If no entities are found,
	// entity.ErrNotFound should be returned instead.
	FindAnalystByUserID(userID string) (*entity.Analyst, error)
}
