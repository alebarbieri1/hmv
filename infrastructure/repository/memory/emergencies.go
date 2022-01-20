package memory

import (
	"flavioltonon/hmv/domain/entity"
	"flavioltonon/hmv/domain/valueobject"
	"sync"
	"time"
)

type Emergency struct {
	ID        string
	UserID    string
	Form      valueobject.EmergencyForm
	CreatedAt time.Time
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
	r.emergencies[emergency.ID] = &Emergency{
		ID:        emergency.ID,
		UserID:    emergency.UserID,
		Form:      emergency.Form,
		CreatedAt: emergency.CreatedAt,
	}
	r.mu.Unlock()
	return nil
}

func (r *EmergenciesRepository) ListEmergencies() ([]*entity.Emergency, error) {
	emergencies := make([]*entity.Emergency, 0, len(r.emergencies))

	r.mu.Lock()

	for _, emergency := range r.emergencies {
		emergencies = append(emergencies, &entity.Emergency{
			ID:        emergency.ID,
			UserID:    emergency.UserID,
			Form:      emergency.Form,
			CreatedAt: emergency.CreatedAt,
		})
	}

	r.mu.Unlock()

	return emergencies, nil
}
