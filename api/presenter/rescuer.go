package presenter

import (
	"flavioltonon/hmv/domain/entity"
)

type Rescuer struct {
	ID        string `json:"_id"`
	UserID    string `json:"user_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewRescuer(e *entity.Rescuer) *Rescuer {
	return &Rescuer{
		ID:        e.ID,
		UserID:    e.UserID,
		CreatedAt: e.CreatedAt.Format("02/01/2006 - 15:04:05h"),
		UpdatedAt: e.UpdatedAt.Format("02/01/2006 - 15:04:05h"),
	}
}
