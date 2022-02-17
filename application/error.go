package application

import "errors"

var (
	ErrInternalError               = errors.New("internal error")
	ErrBasicAuthenticationRequired = errors.New("basic authentication required")
	ErrUserMustBeAPacient          = errors.New("user must be a pacient")
	ErrUserAlreadyIsAPacient       = errors.New("user already is a pacient")
	ErrUsernameAlreadyInUse        = errors.New("username is already in use")
	ErrInvalidUsername             = errors.New("invalid username")
	ErrInvalidPassword             = errors.New("invalid password")
	ErrInvalidUsernameOrPassword   = errors.New("invalid username or password")
)

const (
	UserCreated              = "user created"
	UserAuthenticated        = "user authenticated"
	EmergencyCreated         = "emergency created"
	PacientCreated           = "pacient created"
	PacientUpdated           = "pacient updated"
	FailedToValidateRequest  = "failed to validate request"
	FailedToAuthenticateUser = "failed to authenticate user"
	FailedToCreateEmergency  = "failed to create emergency"
	FailedToListEmergencies  = "failed to list emergencies"
	FailedToCreatePacient    = "failed to create pacient"
	FailedToUpdatePacient    = "failed to update pacient"
	FailedToFindPacient      = "failed to find pacient"
	FailedToCreateUser       = "failed to create user"
)
