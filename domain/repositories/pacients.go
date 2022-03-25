package repositories

import "flavioltonon/hmv/domain/entity"

// PacientsRepository is a repository for entity.Pacient entities
type PacientsRepository interface {
	// CreatePacient stores an entity.Pacient into the repository. If an Pacient with the same ID already
	// exists in the repository, entity.ErrDuplicatedEntry should be returned instead.
	CreatePacient(pacient *entity.Pacient) error

	// FindPacientByID returns an entity.Pacient identified by a given pacientID. If no entities are found,
	// entity.ErrNotFound should be returned instead.
	FindPacientByID(pacientID string) (*entity.Pacient, error)

	// FindPacientByUserID returns an entity.Pacient identified by a given userID. If no entities are found,
	// entity.ErrNotFound should be returned instead.
	FindPacientByUserID(userID string) (*entity.Pacient, error)

	// UpdatePacient updates an Pacient in the repository. If no entities with the same ID as the input entity.Pacient are found,
	// entity.ErrNotFound should be returned instead.
	UpdatePacient(pacient *entity.Pacient) error
}
