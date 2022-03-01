package entity

import (
	"flavioltonon/hmv/domain/valueobject"
	"flavioltonon/hmv/infrastructure/context"
	"net/http"
	"time"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

type User struct {
	ID        string
	Username  string
	Password  string
	Profiles  []valueobject.Profile
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

func NewUserFromRequest(r *http.Request) (*User, error) {
	ctx, err := context.Parse(r.Context())
	if err != nil {
		return nil, err
	}

	user, implements := ctx.Get(context.UserKey).(*User)
	if !implements {
		return nil, ErrNotFound
	}

	return user, nil
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

func (u *User) AddProfile(profile valueobject.Profile) error {
	for _, p := range u.Profiles {
		if p == profile {
			return ErrProfileAlreadySet(profile)
		}
	}

	u.Profiles = append(u.Profiles, profile)
	return nil
}

// IsAnalyst returns true if the User has an Analyst profile
func (u *User) IsAnalyst() bool {
	for _, profile := range u.Profiles {
		if profile == valueobject.AnalystProfile {
			return true
		}
	}

	return false
}

// IsPacient returns true if the User has a Pacient profile
func (u *User) IsPacient() bool {
	for _, profile := range u.Profiles {
		if profile == valueobject.PacientProfile {
			return true
		}
	}

	return false
}
