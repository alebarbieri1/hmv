package usecases

import "flavioltonon/hmv/domain/entity"

// RescuerUsecase defines all the usecases related to a hospital Rescuer
type RescuerUsecase interface {
	// CreateRescuer creates a new entity.Rescuer
	CreateRescuer(userID string) (*entity.Rescuer, error)

	// FindRescuerByID returns an entity.Rescuer with a given rescuerID. If no entities are found, entity.ErrNotFound
	// should be returned instead.
	FindRescuerByID(rescuerID string) (*entity.Rescuer, error)
}
