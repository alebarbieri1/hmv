package entity

import (
	"time"

	"flavioltonon/hmv/domain/valueobject"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

// Emergency defines data about an emergency
type Emergency struct {
	ID         string
	PacientID  string
	Form       valueobject.EmergencyForm
	CreatedAt  time.Time
	UpdatedAt  time.Time
	StatusFlow valueobject.EmergencyStatusFlow
	Status     valueobject.EmergencyStatus
}

// NewEmergency creates a new Emergency
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

// Validate validates an Emergency
func (e *Emergency) Validate() error {
	now := time.Now()

	return ozzo.ValidateStruct(e,
		ozzo.Field(&e.ID, ozzo.Required, is.UUIDv4),
		ozzo.Field(&e.PacientID, ozzo.Required, is.UUIDv4),
		ozzo.Field(&e.CreatedAt, ozzo.Required, ozzo.Max(now)),
		ozzo.Field(&e.UpdatedAt, ozzo.Required, ozzo.Max(now)),
		ozzo.Field(&e.StatusFlow, ozzo.Required),
		ozzo.Field(&e.Status, ozzo.Required),
	)
}

// UpdateForm updates Emergency.Form with a new valueobject.EmergencyForm
func (e *Emergency) UpdateForm(form valueobject.EmergencyForm) error {
	if err := form.Validate(); err != nil {
		return err
	}
	e.Form = form
	return nil
}

// UpdateStatus updates Emergency.Status with a new valueobject.EmergencyStatus. If the status change is not mapped in
// the Emergency.StatusFlow, an error should be returned instead.
func (e *Emergency) UpdateStatus(status valueobject.EmergencyStatus) error {
	if !e.StatusFlow.CanChange(e.Status, status) {
		return ErrInvalidStatusChange(e.Status, status)
	}
	e.Status = status
	return nil
}

// Priority returns the priority level of the Emergency
func (e *Emergency) Priority() valueobject.EmergencyPriority { return e.Form.Priority() }
