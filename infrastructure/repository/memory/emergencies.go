package memory

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
	"sync"
	"time"
)

type Emergency struct {
	ID        string
	PacientID string
	Form      valueobject.EmergencyForm
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    valueobject.EmergencyStatus
}

func NewEmergency(e *entity.Emergency) *Emergency {
	return &Emergency{
		ID:        e.ID,
		PacientID: e.PacientID,
		Form:      e.Form,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		Status:    e.Status,
	}
}

func (e *Emergency) toEntity() *entity.Emergency {
	return &entity.Emergency{
		ID:        e.ID,
		PacientID: e.PacientID,
		Form:      e.Form,
		CreatedAt: e.CreatedAt,
		UpdatedAt: e.UpdatedAt,
		Status:    e.Status,
	}
}

type EmergenciesRepository struct {
	emergencies map[string]*Emergency
	mu          sync.RWMutex
}

func NewEmergenciesRepository() (*EmergenciesRepository, error) {
	return &EmergenciesRepository{emergencies: make(map[string]*Emergency)}, nil
}

func (r *EmergenciesRepository) CreateEmergency(emergency *entity.Emergency) error {
	r.mu.Lock()
	r.emergencies[emergency.ID] = NewEmergency(emergency)
	r.mu.Unlock()
	return nil
}

func (r *EmergenciesRepository) FindEmergencyByID(emergencyID string) (*entity.Emergency, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	emergency, exists := r.emergencies[emergencyID]
	if !exists {
		return nil, entity.ErrNotFound
	}

	return emergency.toEntity(), nil
}

func (r *EmergenciesRepository) ListEmergencies() ([]*entity.Emergency, error) {
	emergencies := make([]*entity.Emergency, 0, len(r.emergencies))

	r.mu.Lock()

	for _, emergency := range r.emergencies {
		emergencies = append(emergencies, emergency.toEntity())
	}

	r.mu.Unlock()

	return emergencies, nil
}

func (r *EmergenciesRepository) ListEmergenciesByStatus(status valueobject.EmergencyStatus) ([]*entity.Emergency, error) {
	emergencies := make([]*entity.Emergency, 0, len(r.emergencies))

	r.mu.Lock()

	for _, emergency := range r.emergencies {
		if emergency.Status != status {
			continue
		}

		emergencies = append(emergencies, emergency.toEntity())
	}

	r.mu.Unlock()

	return emergencies, nil
}

func (r *EmergenciesRepository) ListEmergenciesByPacientID(pacientID string) ([]*entity.Emergency, error) {
	emergencies := make([]*entity.Emergency, 0, len(r.emergencies))

	r.mu.Lock()

	for _, emergency := range r.emergencies {
		if emergency.PacientID != pacientID {
			continue
		}

		emergencies = append(emergencies, emergency.toEntity())
	}

	r.mu.Unlock()

	return emergencies, nil
}

func (r *EmergenciesRepository) UpdateEmergency(emergency *entity.Emergency) error {
	r.mu.Lock()
	r.emergencies[emergency.ID] = NewEmergency(emergency)
	r.mu.Unlock()
	return nil
}
