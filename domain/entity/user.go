package entity

import (
	"time"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type User struct {
	ID        string
	Username  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUser(username, password string) (*User, error) {
	now := time.Now()

	s := &User{
		ID:        uuid.NewString(),
		Username:  username,
		Password:  password,
		CreatedAt: now,
		UpdatedAt: now,
	}

	if err := s.Validate(); err != nil {
		return nil, err
	}

	return s, nil
}

func (u *User) Validate() error {
	now := time.Now()

	return ozzo.ValidateStruct(u,
		ozzo.Field(&u.ID, ozzo.Required, is.UUIDv4),
		ozzo.Field(&u.Username, ozzo.Required, ozzo.Length(0, 64)),
		ozzo.Field(&u.Password, ozzo.Required, ozzo.Length(0, 64)),
		ozzo.Field(&u.CreatedAt, ozzo.Required, ozzo.Max(now)),
		ozzo.Field(&u.UpdatedAt, ozzo.Required, ozzo.Max(now)),
	)
}
