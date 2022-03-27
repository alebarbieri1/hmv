package usecases

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
)

// PacientUsecase defines all the usecases related to a Pacient
type PacientUsecase interface {
	// CreatePacient creates a new entity.Pacient
	CreatePacient(userID string) (*entity.Pacient, error)

	// FindPacientByID returns an entity.Pacient with a given pacientID. If no entities are found, entity.ErrNotFound
	// should be returned instead.
	FindPacientByID(pacientID string) (*entity.Pacient, error)

	// FindPacientByUserID returns an entity.Pacient with a given userID. If no entities are found, entity.ErrNotFound
	// should be returned instead.
	FindPacientByUserID(userID string) (*entity.Pacient, error)

	// UpdateEmergencyContact updates the EmergencyContact of a entity.Pacient with a given pacientID. This action can only
	// be performed by users with an Pacient_ProfileKind.
	UpdateEmergencyContact(userID, pacientID string, emergencyContact valueobject.EmergencyContact) (*entity.Pacient, error)

	// UpdateHealthData updates the HealthData of a entity.Pacient with a given pacientID. This action can only
	// be performed by users with an Pacient_ProfileKind.
	UpdateHealthData(userID, pacientID string, healthData valueobject.HealthData) (*entity.Pacient, error)

	// UpdateLocationData updates the LocationData of a entity.Pacient with a given pacientID. This action can only
	// be performed by users with an Pacient_ProfileKind.
	UpdateLocationData(userID, pacientID string, locationData valueobject.LocationData) (*entity.Pacient, error)
}
