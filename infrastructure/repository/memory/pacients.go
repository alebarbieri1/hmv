package memory

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
	"sync"
	"time"
)

type Pacient struct {
	ID               string
	UserID           string
	EmergencyContact valueobject.EmergencyContact
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func NewPacient(e *entity.Pacient) *Pacient {
	return &Pacient{
		ID:               e.ID,
		UserID:           e.UserID,
		EmergencyContact: e.EmergencyContact,
		CreatedAt:        e.CreatedAt,
		UpdatedAt:        e.UpdatedAt,
	}
}

func (u *Pacient) toEntity() *entity.Pacient {
	return &entity.Pacient{
		ID:               u.ID,
		UserID:           u.UserID,
		EmergencyContact: u.EmergencyContact,
		CreatedAt:        u.CreatedAt,
		UpdatedAt:        u.UpdatedAt,
	}
}

type PacientsRepository struct {
	pacients map[string]*Pacient
	mu       sync.RWMutex
}

func NewPacientsRepository() (*PacientsRepository, error) {
	return &PacientsRepository{pacients: make(map[string]*Pacient)}, nil
}

func (r *PacientsRepository) CreatePacient(pacient *entity.Pacient) error {
	r.mu.Lock()
	r.pacients[pacient.ID] = NewPacient(pacient)
	r.mu.Unlock()
	return nil
}

func (r *PacientsRepository) FindPacientByID(pacientID string) (*entity.Pacient, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for _, pacient := range r.pacients {
		if pacient.ID == pacientID {
			return pacient.toEntity(), nil
		}
	}

	return nil, entity.ErrNotFound
}

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

func (r *PacientsRepository) UpdatePacient(pacient *entity.Pacient) error {
	r.mu.Lock()
	r.pacients[pacient.ID] = NewPacient(pacient)
	r.mu.Unlock()
	return nil
}
