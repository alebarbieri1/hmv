package entity

import (
	"flavioltonon/hmv/domain/valueobject"
	"time"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

// Pacient defines data about a pacient
type Pacient struct {
	ID        string
	UserID    string
	Data      valueobject.PacientData
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Validate validates a Pacient
func (p *Pacient) Validate() error {
	now := time.Now()

	return ozzo.ValidateStruct(p,
		ozzo.Field(&p.ID, ozzo.Required, is.UUIDv4),
		ozzo.Field(&p.UserID, ozzo.Required, is.UUIDv4),
		ozzo.Field(&p.CreatedAt, ozzo.Required, ozzo.Max(now)),
		ozzo.Field(&p.UpdatedAt, ozzo.Required, ozzo.Max(now)),
	)
}

// UpdateLocationData updates the pacient's emergency contact data
func (p *Pacient) UpdateLocationData(data valueobject.LocationData) error {
	if err := data.Validate(); err != nil {
		return err
	}
	p.Data.Location = data
	p.UpdatedAt = time.Now()
	return nil
}

// UpdateEmergencyContact updates the pacient's emergency contact data
func (p *Pacient) UpdateEmergencyContact(emergencyContact valueobject.EmergencyContact) error {
	if err := emergencyContact.Validate(); err != nil {
		return err
	}
	p.Data.EmergencyContact = emergencyContact
	p.UpdatedAt = time.Now()
	return nil
}

// UpdateHealthData updates the pacient's emergency contact data
func (p *Pacient) UpdateHealthData(data valueobject.HealthData) error {
	p.Data.Health = data
	p.UpdatedAt = time.Now()
	return nil
}
