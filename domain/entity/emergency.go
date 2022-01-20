package entity

import (
	"time"

	"flavioltonon/hmv/domain/valueobject"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type Emergency struct {
	ID        string
	UserID    string
	Form      valueobject.EmergencyForm
	CreatedAt time.Time
}

func NewEmergency(userID string) (*Emergency, error) {
	e := &Emergency{
		ID:        uuid.NewString(),
		UserID:    userID,
		CreatedAt: time.Now(),
	}

	if err := e.Validate(); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Emergency) Validate() error {
	return ozzo.ValidateStruct(e,
		ozzo.Field(&e.UserID, ozzo.Required, is.UUIDv4),
	)
}

func (e *Emergency) UpdateForm(form valueobject.EmergencyForm) { e.Form = form }

func (e *Emergency) Priority() valueobject.EmergencyPriority { return e.Form.Priority() }
