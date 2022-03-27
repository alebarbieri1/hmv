package presenter

import (
	"flavioltonon/hmv/domain/entity"
)

// User is a entity.User presenter
type User struct {
	ID          string `json:"_id"`
	Name        string `json:"name"`
	ProfileKind string `json:"profile_kind"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// NewUser returns a presentation for a User
func NewUser(e *entity.User) *User {
	return &User{
		ID:          e.ID,
		Name:        e.Data.Name,
		ProfileKind: e.ProfileKind.String(),
		CreatedAt:   e.CreatedAt.Format("02/01/2006 - 15:04:05h"),
		UpdatedAt:   e.UpdatedAt.Format("02/01/2006 - 15:04:05h"),
	}
}

// NewUsers returns a presentation for a set of Users
func NewUsers(es []*entity.User) []*User {
	users := make([]*User, 0, len(es))

	for _, e := range es {
		users = append(users, NewUser(e))
	}

	return users
}
