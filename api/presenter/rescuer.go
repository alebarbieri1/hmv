package presenter

import (
	"flavioltonon/hmv/domain/entity"
)

// Rescuer is a entity.Rescuer presenter
type Rescuer struct {
	ID        string `json:"_id"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

// NewRescuer returns a presentation for a Rescuer
func NewRescuer(e *entity.Rescuer) *Rescuer {
	return &Rescuer{
		ID:        e.ID,
		UserID:    e.UserID,
		CreatedAt: e.CreatedAt.Format("02/01/2006 - 15:04:05h"),
		UpdatedAt: e.UpdatedAt.Format("02/01/2006 - 15:04:05h"),
	}
}
