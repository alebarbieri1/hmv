package application

import "errors"

var (
	ErrInternalError               = errors.New("internal error")
	ErrBasicAuthenticationRequired = errors.New("basic authentication required")
	ErrUserMustBeAnAnalyst         = errors.New("user must be an analyst")
	ErrUserAlreadyIsAnAnalyst      = errors.New("user already is an analyst")
	ErrUserMustBeAPacient          = errors.New("user must be a pacient")
	ErrUserAlreadyIsAPacient       = errors.New("user already is a pacient")
	ErrInvalidUserProfiles         = errors.New("invalid user profiles")
	ErrUsernameAlreadyInUse        = errors.New("username is already in use")
	ErrInvalidUsername             = errors.New("invalid username")
	ErrInvalidPassword             = errors.New("invalid password")
	ErrInvalidUsernameOrPassword   = errors.New("invalid username or password")
)

const (
	UserFound                = "user found"
	UserCreated              = "user created"
	UserCanBeUpdated         = "user can be updated"
	UserUpdated              = "user updated"
	UserAuthenticated        = "user authenticated"
	EmergencyCreated         = "emergency created"
	AnalystCreated           = "analyst created"
	PacientCreated           = "pacient created"
	PacientUpdated           = "pacient updated"
	FailedToValidateRequest  = "failed to validate request"
	FailedToAuthenticateUser = "failed to authenticate user"
	FailedToCreateEmergency  = "failed to create emergency"
	FailedToListEmergencies  = "failed to list emergencies"
	FailedToCreateAnalyst    = "failed to create analyst"
	FailedToFindAnalyst      = "failed to find analyst"
	FailedToCreatePacient    = "failed to create pacient"
	FailedToUpdatePacient    = "failed to update pacient"
	FailedToFindPacient      = "failed to find pacient"
	FailedToFindUser         = "failed to find user"
	FailedToCreateUser       = "failed to create user"
	FailedToUpdateUser       = "failed to update user"
)
