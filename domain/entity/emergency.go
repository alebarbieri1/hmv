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
	PacientID string
	Form      valueobject.EmergencyForm
	CreatedAt time.Time
	UpdatedAt time.Time
	Status    valueobject.EmergencyStatus
}

func NewEmergency(pacientID string) (*Emergency, error) {
	now := time.Now()

	e := &Emergency{
		ID:        uuid.NewString(),
		PacientID: pacientID,
		CreatedAt: now,
		UpdatedAt: now,
		Status:    valueobject.Triage_EmergencyStatus,
	}

	if err := e.Validate(); err != nil {
		return nil, err
	}

	return e, nil
}

func (e *Emergency) Validate() error {
	now := time.Now()

	return ozzo.ValidateStruct(e,
		ozzo.Field(&e.ID, ozzo.Required, is.UUIDv4),
		ozzo.Field(&e.PacientID, ozzo.Required, is.UUIDv4),
		ozzo.Field(&e.CreatedAt, ozzo.Required, ozzo.Max(now)),
		ozzo.Field(&e.UpdatedAt, ozzo.Required, ozzo.Max(now)),
		ozzo.Field(&e.Status, ozzo.Required),
	)
}

func (e *Emergency) UpdateForm(form valueobject.EmergencyForm) { e.Form = form }

func (e *Emergency) UpdateStatus(status valueobject.EmergencyStatus) { e.Status = status }

func (e *Emergency) Priority() valueobject.EmergencyPriority { return e.Form.Priority() }
