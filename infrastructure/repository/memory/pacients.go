package memory

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
	"sync"
	"time"
)

// Pacient is a representation of entity.Pacient in the repository
type Pacient struct {
	ID               string
	UserID           string
	EmergencyContact valueobject.EmergencyContact
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// NewPacient creates a new Pacient
func NewPacient(e *entity.Pacient) *Pacient {
	return &Pacient{
		ID:               e.ID,
		UserID:           e.UserID,
		EmergencyContact: e.EmergencyContact,
		CreatedAt:        e.CreatedAt,
		UpdatedAt:        e.UpdatedAt,
	}
}

// toEntity transforms an Pacient into a entity.Pacient
func (u *Pacient) toEntity() *entity.Pacient {
	return &entity.Pacient{
		ID:               u.ID,
		UserID:           u.UserID,
		EmergencyContact: u.EmergencyContact,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
	}
}

// PacientsRepository is a repository for entity.Pacient entities
type PacientsRepository struct {
	pacients map[string]*Pacient
	mu       sync.RWMutex
}

// NewPacientsRepository creates a new PacientsRepository
func NewPacientsRepository() (*PacientsRepository, error) {
	return &PacientsRepository{pacients: make(map[string]*Pacient)}, nil
}

// CreatePacient stores an entity.Pacient into the repository. If an Pacient with the same ID already
// exists in the repository, entity.ErrDuplicatedEntry should be returned instead.
func (r *PacientsRepository) CreatePacient(pacient *entity.Pacient) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.pacients[pacient.ID]; exists {
		return entity.ErrDuplicatedEntry
	}

	r.pacients[pacient.ID] = NewPacient(pacient)
	return nil
}

// FindPacientByID returns an entity.Pacient identified by a given pacientID. If no entities are found,
// entity.ErrNotFound should be returned instead.
func (r *PacientsRepository) FindPacientByID(pacientID string) (*entity.Pacient, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	pacient, exists := r.pacients[pacientID]
	if !exists {
		return nil, entity.ErrNotFound
	}

	return pacient.toEntity(), nil
}

// FindPacientByUserID returns an entity.Pacient identified by a given userID. If no entities are found,
// entity.ErrNotFound should be returned instead.
func (r *PacientsRepository) FindPacientByUserID(userID string) (*entity.Pacient, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, pacient := range r.pacients {
		if pacient.UserID == userID {
			return pacient.toEntity(), nil
		}
	}

	return nil, entity.ErrNotFound
}

// UpdatePacient updates an Pacient in the repository. If no entities with the same ID as the input entity.Pacient are found,
// entity.ErrNotFound should be returned instead.
func (r *PacientsRepository) UpdatePacient(pacient *entity.Pacient) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.pacients[pacient.ID]; !exists {
		return entity.ErrNotFound
	}

	r.pacients[pacient.ID] = NewPacient(pacient)
	return nil
}
