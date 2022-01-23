package presenter

import (
	"flavioltonon/hmv/domain/entity"
)

type Emergency struct {
	ID        string         `json:"_id"`
	PacientID string         `json:"pacient_id"`
	Form      *EmergencyForm `json:"form,omitempty"`
	CreatedAt string         `json:"created_at"`
}

func NewEmergency(e *entity.Emergency) *Emergency {
	return &Emergency{
		ID:        e.ID,
		PacientID: e.PacientID,
		Form:      NewEmergencyForm(e.Form),
		CreatedAt: e.CreatedAt.Format("02/01/2006 - 15:04:05h"),
	}
}

type Emergencies []*Emergency

func NewEmergencies(es []*entity.Emergency) Emergencies {
	emergencies := make(Emergencies, 0, len(es))

	for _, emergency := range es {
		emergencies = append(emergencies, NewEmergency(emergency))
	}

	return emergencies
}
