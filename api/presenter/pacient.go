package presenter

import (
	"flavioltonon/hmv/domain/entity"
)

// Pacient is a entity.Pacient presenter
type Pacient struct {
	ID               string            `json:"_id"`
	UserID           string            `json:"user_id"`
	Name             string            `json:"name"`
	BirthDate        string            `json:"birth_date"`
	Location         *LocationData     `json:"location"`
	EmergencyContact *EmergencyContact `json:"emergency_contact"`
	Health           *HealthData       `json:"health"`
	CreatedAt        string            `json:"created_at"`
	UpdatedAt        string            `json:"updated_at"`
}

// NewPacient returns a presentation for a Pacient
func NewPacient(e *entity.Pacient) *Pacient {
	pacient := &Pacient{
		ID:               e.ID,
		UserID:           e.UserID,
		Name:             e.Data.Name,
		Location:         NewLocationData(e.Data.Location),
		EmergencyContact: NewEmergencyContact(e.Data.EmergencyContact),
		Health:           NewHealthData(e.Data.Health),
		CreatedAt:        e.CreatedAt.Format("02/01/2006 - 15:04:05h"),
		UpdatedAt:        e.UpdatedAt.Format("02/01/2006 - 15:04:05h"),
	}

	if !e.Data.BirthDate.IsZero() {
		pacient.BirthDate = e.Data.BirthDate.Format("02/01/2006 - 15:04:05h")

	}

	return pacient
}
