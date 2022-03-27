package memory

import (
	"flavioltonon/hmv/domain/entity"
	"sync"
	"time"
)

// Analyst is a representation of entity.Analyst in the repository
type Analyst struct {
	ID        string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewAnalyst creates a new Analyst
func NewAnalyst(e *entity.Analyst) *Analyst {
	return &Analyst{
		ID:        e.ID,
		UserID:    e.UserID,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

// toEntity transforms an Analyst into a entity.Analyst
func (u *Analyst) toEntity() *entity.Analyst {
	return &entity.Analyst{
		ID:        u.ID,
		UserID:    u.UserID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// AnalystsRepository is a repository for entity.Analyst entities
type AnalystsRepository struct {
	analysts map[string]*Analyst
	mu       sync.RWMutex
}

// NewAnalystsRepository creates a new AnalystsRepository
func NewAnalystsRepository() *AnalystsRepository {
	return &AnalystsRepository{analysts: make(map[string]*Analyst)}
}

// CreateAnalyst stores an entity.Analyst into the repository. If an Analyst with the same ID already
// exists in the repository, entity.ErrDuplicatedEntry should be returned instead.
func (r *AnalystsRepository) CreateAnalyst(analyst *entity.Analyst) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.analysts[analyst.ID]; exists {
		return entity.ErrDuplicatedEntry
	}

	r.analysts[analyst.ID] = NewAnalyst(analyst)
	return nil
}

// FindAnalystByID returns an entity.Analyst identified by a given analystID. If no entities are found,
// entity.ErrNotFound should be returned instead.
func (r *AnalystsRepository) FindAnalystByID(analystID string) (*entity.Analyst, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	analyst, exists := r.analysts[analystID]
	if !exists {
		return nil, entity.ErrNotFound
	}

	return analyst.toEntity(), nil
}

// FindAnalystByUserID returns an entity.Analyst identified by a given userID. If no entities are found,
// entity.ErrNotFound should be returned instead.
func (r *AnalystsRepository) FindAnalystByUserID(userID string) (*entity.Analyst, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, analyst := range r.analysts {
		if analyst.UserID == userID {
			return analyst.toEntity(), nil
		}
	}

	return nil, entity.ErrNotFound
}
