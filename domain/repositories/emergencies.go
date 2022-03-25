package repositories

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
)

// EmergenciesRepository is a repository for entity.Emergency entities
type EmergenciesRepository interface {
	// CreateEmergency stores an entity.Emergency into the repository. If an Emergency with the same ID already
	// exists in the repository, entity.ErrDuplicatedEntry should be returned instead.
	CreateEmergency(emergency *entity.Emergency) error

	// FindEmergencyByID returns an entity.Emergency identified by a given emergencyID. If no entities are found,
	// entity.ErrNotFound should be returned instead.
	FindEmergencyByID(emergencyID string) (*entity.Emergency, error)

	// ListEmergencies returns all the entity.Emergency of the repository
	ListEmergencies() ([]*entity.Emergency, error)

	// ListEmergenciesByStatus returns all the entity.Emergency of the repository that currently have a given valueobject.EmergencyStatus
	ListEmergenciesByStatus(status valueobject.EmergencyStatus) ([]*entity.Emergency, error)

	// ListEmergenciesByPacientID returns all the entity.Emergency of the repository that are related to a entity.Pacient with a given pacientID
	ListEmergenciesByPacientID(pacientID string) ([]*entity.Emergency, error)

	// UpdateEmergency updates an Emergency in the repository. If no entities with the same ID as the input entity.Emergency are found,
	// entity.ErrNotFound should be returned instead.
	UpdateEmergency(emergency *entity.Emergency) error
}
