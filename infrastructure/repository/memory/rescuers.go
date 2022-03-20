package memory

import (
	"flavioltonon/hmv/domain/entity"
	"sync"
	"time"
)

type Rescuer struct {
	ID        string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewRescuer(e *entity.Rescuer) *Rescuer {
	return &Rescuer{
		ID:        e.ID,
		UserID:    e.UserID,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
	}
}

func (u *Rescuer) toEntity() *entity.Rescuer {
	return &entity.Rescuer{
		ID:        u.ID,
		UserID:    u.UserID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

type RescuersRepository struct {
	rescuers map[string]*Rescuer
	mu       sync.RWMutex
}

func NewRescuersRepository() (*RescuersRepository, error) {
	return &RescuersRepository{rescuers: make(map[string]*Rescuer)}, nil
}

func (r *RescuersRepository) CreateRescuer(rescuer *entity.Rescuer) error {
	r.mu.Lock()
	r.rescuers[rescuer.ID] = NewRescuer(rescuer)
	r.mu.Unlock()
	return nil
}

func (r *RescuersRepository) FindRescuerByID(rescuerID string) (*entity.Rescuer, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	rescuer, exists := r.rescuers[rescuerID]
	if !exists {
		return nil, entity.ErrNotFound
	}

	return rescuer.toEntity(), nil
}

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
