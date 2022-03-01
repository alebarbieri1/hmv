package memory

import (
	"flavioltonon/hmv/domain/entity"
	"sync"
	"time"
)

type Analyst struct {
	ID        string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewAnalyst(e *entity.Analyst) *Analyst {
	return &Analyst{
		ID:        e.ID,
		UserID:    e.UserID,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func (u *Analyst) toEntity() *entity.Analyst {
	return &entity.Analyst{
		ID:        u.ID,
		UserID:    u.UserID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type AnalystsRepository struct {
	analysts map[string]*Analyst
	mu       sync.RWMutex
}

func NewAnalystsRepository() (*AnalystsRepository, error) {
	return &AnalystsRepository{analysts: make(map[string]*Analyst)}, nil
}

func (r *AnalystsRepository) CreateAnalyst(analyst *entity.Analyst) error {
	r.mu.Lock()
	r.analysts[analyst.ID] = NewAnalyst(analyst)
	r.mu.Unlock()
	return nil
}

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
