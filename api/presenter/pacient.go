package presenter

import (
	"flavioltonon/hmv/domain/entity"
)

// Pacient is a entity.Pacient presenter
type Pacient struct {
	ID               string            `json:"_id"`
	UserID           string            `json:"user_id"`
	EmergencyContact *EmergencyContact `json:"emergency_contact"`
	CreatedAt        string            `json:"created_at"`
	UpdatedAt        string            `json:"updated_at"`
}

// NewPacient returns a presentation for a Pacient
func NewPacient(e *entity.Pacient) *Pacient {
	return &Pacient{
		ID:               e.ID,
		UserID:           e.UserID,
		EmergencyContact: NewEmergencyContact(e.EmergencyContact),
		CreatedAt:        e.CreatedAt.Format("02/01/2006 - 15:04:05h"),
		UpdatedAt:        e.UpdatedAt.Format("02/01/2006 - 15:04:05h"),
	}
}
