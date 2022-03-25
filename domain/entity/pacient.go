package entity

import (
	"flavioltonon/hmv/domain/valueobject"
	"time"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

// Pacient defines data about a pacient
type Pacient struct {
	ID               string
	UserID           string
	EmergencyContact valueobject.EmergencyContact
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

// NewPacient creates a new Pacient
func NewPacient(userID string) (*Pacient, error) {
	now := time.Now()

	s := &Pacient{
		ID:        uuid.NewString(),
		UserID:    userID,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.Validate(); err != nil {
		return nil, err
	}

	return s, nil
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

// UpdateEmergencyContact updates the pacient's emergency contact data
func (p *Pacient) UpdateEmergencyContact(emergencyContact valueobject.EmergencyContact) error {
	if err := emergencyContact.Validate(); err != nil {
		return err
	}
	p.EmergencyContact = emergencyContact
	p.UpdatedAt = time.Now()
	return nil
}
