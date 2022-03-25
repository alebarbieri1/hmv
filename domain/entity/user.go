package entity

import (
	"context"
	"flavioltonon/hmv/domain/valueobject"
	internalContext "flavioltonon/hmv/infrastructure/context"
	"net/http"
	"time"

	ozzo "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
)

// User is one of the application's users
type User struct {
	ID          string
	Username    string
	Password    string
	ProfileKind valueobject.ProfileKind
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// NewUser creates a new User with a given username/password pair
func NewUser(username, password string) (*User, error) {
	now := time.Now()

	s := &User{
		ID:          uuid.NewString(),
		Username:    username,
		Password:    password,
		ProfileKind: valueobject.Undefined_ProfileKind,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.Validate(); err != nil {
		return nil, err
	}

	return s, nil
}

// NewUserFromRequest recovers a User from a given http.Request
func NewUserFromRequest(r *http.Request) (*User, error) {
	return NewUserFromContext(r.Context())
}

// NewUserFromContext recovers a User from a given context.Context
func NewUserFromContext(ctx context.Context) (*User, error) {
	ictx, err := internalContext.Parse(ctx)
	if err != nil {
		return nil, err
	}

	user, implements := ictx.Get(internalContext.UserKey).(*User)
	if !implements {
		return nil, ErrNotFound
	}

	return user, nil
}

// Validate validates a User
func (u *User) Validate() error {
	now := time.Now()

	return ozzo.ValidateStruct(u,
		ozzo.Field(&u.ID, ozzo.Required, is.UUIDv4),
		ozzo.Field(&u.Username, ozzo.Required, ozzo.Length(0, 64)),
		ozzo.Field(&u.Password, ozzo.Required, ozzo.Length(0, 64)),
		ozzo.Field(&u.ProfileKind, ozzo.Required),
		ozzo.Field(&u.CreatedAt, ozzo.Required, ozzo.Max(now)),
		ozzo.Field(&u.UpdatedAt, ozzo.Required, ozzo.Max(now)),
	)
}

// SetProfileKind sets the User with a ProfileKind. If the User already has a ProfileKind, an error should be returned instead.
func (u *User) SetProfileKind(profileKind valueobject.ProfileKind) error {
	if u.HasProfileKind() {
		return ErrProfileKindAlreadySet(u.ProfileKind)
	}

	u.ProfileKind = profileKind
	return nil
}

// HasProfileKind returns true if the user already has a profile set
func (u *User) HasProfileKind() bool { return u.ProfileKind != valueobject.Undefined_ProfileKind }

// IsAnalyst returns true if the User has an Analyst profile
func (u *User) IsAnalyst() bool { return u.ProfileKind == valueobject.Analyst_ProfileKind }

// IsPacient returns true if the User has a Pacient profile
func (u *User) IsPacient() bool { return u.ProfileKind == valueobject.Pacient_ProfileKind }

// IsRescuer returns true if the User has a Rescuer profile
func (u *User) IsRescuer() bool { return u.ProfileKind == valueobject.Rescuer_ProfileKind }
