package repositories

import "flavioltonon/hmv/domain/entity"

// RescuersRepository is a repository for entity.Rescuer entities
type RescuersRepository interface {
	// CreateRescuer stores an entity.Rescuer into the repository. If an Rescuer with the same ID already
	// exists in the repository, entity.ErrDuplicatedEntry should be returned instead.
	CreateRescuer(pacient *entity.Rescuer) error

	// FindRescuerByID returns an entity.Rescuer identified by a given rescuerID. If no entities are found,
	// entity.ErrNotFound should be returned instead.
	FindRescuerByID(rescuerID string) (*entity.Rescuer, error)

	// FindRescuerByUserID returns an entity.Rescuer identified by a given userID. If no entities are found,
	// entity.ErrNotFound should be returned instead.
	FindRescuerByUserID(userID string) (*entity.Rescuer, error)
}
