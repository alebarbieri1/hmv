package entity

import (
	"time"

	"flavioltonon/hmv/domain/valueobject"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type Emergency struct {
	ID         string
	PacientID  string
	Form       valueobject.EmergencyForm
	CreatedAt  time.Time
	UpdatedAt  time.Time
	StatusFlow valueobject.EmergencyStatusFlow
	Status     valueobject.EmergencyStatus
}

func NewEmergency(pacientID string) (*Emergency, error) {
	now := time.Now()

	e := &Emergency{
		ID:         uuid.NewString(),
		PacientID:  pacientID,
		CreatedAt:  now,
		UpdatedAt:  now,
		StatusFlow: valueobject.DefaultEmergencyStatusFlow,
		Status:     valueobject.Triage_EmergencyStatus,
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
		ozzo.Field(&e.Form),
		ozzo.Field(&e.CreatedAt, ozzo.Required, ozzo.Max(now)),
		ozzo.Field(&e.UpdatedAt, ozzo.Required, ozzo.Max(now)),
		ozzo.Field(&e.Status, ozzo.Required),
	)
}

func (e *Emergency) UpdateForm(form valueobject.EmergencyForm) error {
	e.Form = form
	return e.Validate()
}

func (e *Emergency) UpdateStatus(status valueobject.EmergencyStatus) error {
	if !e.StatusFlow.CanChange(e.Status, status) {
		return ErrInvalidStatusChange(e.Status, status)
	}
	e.Status = status
	return e.Validate()
}

func (e *Emergency) Priority() valueobject.EmergencyPriority { return e.Form.Priority() }
