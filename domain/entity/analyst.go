package entity

import (
	"time"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

// Analyst is one of the Hospital users responsible for managing emergencies and allocating resources for
// them, such as ambulances, doctors, and others.
type Analyst struct {
	ID        string
	UserID    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewAnalyst creates a new Analyst
func NewAnalyst(userID string) (*Analyst, error) {
	now := time.Now()

	s := &Analyst{
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

// Validate validates an Analyst
func (p *Analyst) Validate() error {
	now := time.Now()

	return ozzo.ValidateStruct(p,
		ozzo.Field(&p.ID, ozzo.Required, is.UUIDv4),
		ozzo.Field(&p.UserID, ozzo.Required, is.UUIDv4),
		ozzo.Field(&p.CreatedAt, ozzo.Max(now)),
		ozzo.Field(&p.UpdatedAt, ozzo.Max(now)),
	)
}
