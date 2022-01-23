package entity

import (
	"time"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type Pacient struct {
	ID        string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

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

func (p *Pacient) Validate() error {
	now := time.Now()

	return ozzo.ValidateStruct(p,
		ozzo.Field(&p.ID, ozzo.Required, is.UUIDv4),
		ozzo.Field(&p.UserID, ozzo.Required, is.UUIDv4),
		ozzo.Field(&p.CreatedAt, ozzo.Required, ozzo.Max(now)),
		ozzo.Field(&p.UpdatedAt, ozzo.Required, ozzo.Max(now)),
	)
}

func (p *Pacient) CreateEmergency() (*Emergency, error) {
	return NewEmergency(p.ID)
}
