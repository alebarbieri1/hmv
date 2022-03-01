package presenter

import (
	"flavioltonon/hmv/domain/entity"
)

type User struct {
	ID          string `json:"_id"`
	ProfileKind string `json:"profile_kind"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func NewUser(e *entity.User) *User {
	return &User{
		ID:          e.ID,
		ProfileKind: e.ProfileKind.String(),
		CreatedAt:   e.CreatedAt.Format("02/01/2006 - 15:04:05h"),
		UpdatedAt:   e.UpdatedAt.Format("02/01/2006 - 15:04:05h"),
	}
}
