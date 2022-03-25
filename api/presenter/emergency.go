package presenter

import (
	"flavioltonon/hmv/domain/entity"
)

// Emergency is a entity.Emergency presenter
type Emergency struct {
	ID        string         `json:"_id"`
	PacientID string         `json:"pacient_id"`
	Form      *EmergencyForm `json:"form,omitempty"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
	Status    string         `json:"status"`
}

// NewEmergency returns a presentation for a Emergency
func NewEmergency(e *entity.Emergency) *Emergency {
	return &Emergency{
		ID:        e.ID,
		PacientID: e.PacientID,
		Form:      NewEmergencyForm(e.Form),
		CreatedAt: e.CreatedAt.Format("02/01/2006 - 15:04:05h"),
		UpdatedAt: e.UpdatedAt.Format("02/01/2006 - 15:04:05h"),
		Status:    e.Status.String(),
	}
}

// Emergencies is a set of Emergency presenters
type Emergencies []*Emergency

// NewEmergencies returns a presentation for Emergencies
func NewEmergencies(es []*entity.Emergency) Emergencies {
	emergencies := make(Emergencies, 0, len(es))

	for _, emergency := range es {
		emergencies = append(emergencies, NewEmergency(emergency))
	}

	return emergencies
}
