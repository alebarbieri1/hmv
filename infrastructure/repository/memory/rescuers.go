package memory

import (
	"flavioltonon/hmv/domain/entity"
	"sync"
	"time"
)

// Rescuer is a representation of entity.Rescuer in the repository
type Rescuer struct {
	ID        string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewRescuer creates a new Rescuer
func NewRescuer(e *entity.Rescuer) *Rescuer {
	return &Rescuer{
		ID:        e.ID,
		UserID:    e.UserID,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

// toEntity transforms an Rescuer into a entity.Rescuer
func (u *Rescuer) toEntity() *entity.Rescuer {
	return &entity.Rescuer{
		ID:        u.ID,
		UserID:    u.UserID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// RescuersRepository is a repository for entity.Rescuer entities
type RescuersRepository struct {
	rescuers map[string]*Rescuer
	mu       sync.RWMutex
}

// NewRescuersRepository creates a new RescuersRepository
func NewRescuersRepository() *RescuersRepository {
	return &RescuersRepository{rescuers: make(map[string]*Rescuer)}
}

// CreateRescuer stores an entity.Rescuer into the repository. If an Rescuer with the same ID already
// exists in the repository, entity.ErrDuplicatedEntry should be returned instead.
func (r *RescuersRepository) CreateRescuer(rescuer *entity.Rescuer) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.rescuers[rescuer.ID]; exists {
		return entity.ErrDuplicatedEntry
	}

	r.rescuers[rescuer.ID] = NewRescuer(rescuer)
	return nil
}

// FindRescuerByID returns an entity.Rescuer identified by a given rescuerID. If no entities are found,
// entity.ErrNotFound should be returned instead.
func (r *RescuersRepository) FindRescuerByID(rescuerID string) (*entity.Rescuer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	rescuer, exists := r.rescuers[rescuerID]
	if !exists {
		return nil, entity.ErrNotFound
	}

	return rescuer.toEntity(), nil
}

// FindRescuerByUserID returns an entity.Rescuer identified by a given userID. If no entities are found,
// entity.ErrNotFound should be returned instead.
func (r *RescuersRepository) FindRescuerByUserID(userID string) (*entity.Rescuer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, rescuer := range r.rescuers {
		if rescuer.UserID == userID {
			return rescuer.toEntity(), nil
		}
	}

	return nil, entity.ErrNotFound
}
