package memory

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
	"sync"
	"time"
)

// Emergency is a representation of entity.Emergency in the repository
type Emergency struct {
	ID         string
	PacientID  string
	Form       valueobject.EmergencyForm
	CreatedAt  time.Time
	UpdatedAt  time.Time
	StatusFlow valueobject.EmergencyStatusFlow
	Status     valueobject.EmergencyStatus
}

// NewEmergency creates a new Emergency
func NewEmergency(e *entity.Emergency) *Emergency {
	return &Emergency{
		ID:         e.ID,
		PacientID:  e.PacientID,
		Form:       e.Form,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
		StatusFlow: e.StatusFlow,
		Status:     e.Status,
	}
}

// toEntity transforms an Emergency into a entity.Emergency
func (e *Emergency) toEntity() *entity.Emergency {
	return &entity.Emergency{
		ID:         e.ID,
		PacientID:  e.PacientID,
		Form:       e.Form,
		CreatedAt:  e.CreatedAt,
		UpdatedAt:  e.UpdatedAt,
		StatusFlow: e.StatusFlow,
		Status:     e.Status,
	}
}

// EmergenciesRepository is a repository for entity.Emergency entities
type EmergenciesRepository struct {
	emergencies map[string]*Emergency
	mu          sync.RWMutex
}

// NewEmergenciesRepository creates a new EmergenciesRepository
func NewEmergenciesRepository() (*EmergenciesRepository, error) {
	return &EmergenciesRepository{emergencies: make(map[string]*Emergency)}, nil
}

// CreateEmergency stores an entity.Emergency into the repository. If an Emergency with the same ID already
// exists in the repository, entity.ErrDuplicatedEntry should be returned instead.
func (r *EmergenciesRepository) CreateEmergency(emergency *entity.Emergency) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.emergencies[emergency.ID]; exists {
		return entity.ErrDuplicatedEntry
	}

	r.emergencies[emergency.ID] = NewEmergency(emergency)
	return nil
}

// FindEmergencyByID returns an entity.Emergency identified by a given emergencyID. If no entities are found,
// entity.ErrNotFound should be returned instead.
func (r *EmergenciesRepository) FindEmergencyByID(emergencyID string) (*entity.Emergency, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	emergency, exists := r.emergencies[emergencyID]
	if !exists {
		return nil, entity.ErrNotFound
	}

	return emergency.toEntity(), nil
}

// ListEmergencies returns all the entity.Emergency of the repository
func (r *EmergenciesRepository) ListEmergencies() ([]*entity.Emergency, error) {
	emergencies := make([]*entity.Emergency, 0, len(r.emergencies))

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, emergency := range r.emergencies {
		emergencies = append(emergencies, emergency.toEntity())
	}

	return emergencies, nil
}

// ListEmergenciesByStatus returns all the entity.Emergency of the repository that currently have a given valueobject.EmergencyStatus
func (r *EmergenciesRepository) ListEmergenciesByStatus(status valueobject.EmergencyStatus) ([]*entity.Emergency, error) {
	var emergencies []*entity.Emergency

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, emergency := range r.emergencies {
		if emergency.Status != status {
			continue
		}

		emergencies = append(emergencies, emergency.toEntity())
	}

	return emergencies, nil
}

// ListEmergenciesByPacientID returns all the entity.Emergency of the repository that are related to a entity.Pacient with a given pacientID
func (r *EmergenciesRepository) ListEmergenciesByPacientID(pacientID string) ([]*entity.Emergency, error) {
	var emergencies []*entity.Emergency

	r.mu.Lock()
	defer r.mu.Unlock()

	for _, emergency := range r.emergencies {
		if emergency.PacientID != pacientID {
			continue
		}

		emergencies = append(emergencies, emergency.toEntity())
	}

	return emergencies, nil
}

// UpdateEmergency updates an Emergency in the repository. If no entities with the same ID as the input entity.Emergency are found,
// entity.ErrNotFound should be returned instead.
func (r *EmergenciesRepository) UpdateEmergency(emergency *entity.Emergency) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.emergencies[emergency.ID]; !exists {
		return entity.ErrNotFound
	}

	r.emergencies[emergency.ID] = NewEmergency(emergency)
	return nil
}
