package application

import "errors"

var (
	ErrBasicAuthenticationRequired = errors.New("basic authentication required")
	ErrUserMustBeAPacient          = errors.New("user must be a pacient")
	ErrUserAlreadyIsAPacient       = errors.New("user already is a pacient")
	ErrUsernameAlreadyInUse        = errors.New("username is already in use")
	ErrInvalidUsername             = errors.New("invalid username")
	ErrInvalidPassword             = errors.New("invalid password")
	ErrInvalidUsernameOrPassword   = errors.New("invalid username or password")
)

const (
	ErrMsgFailedToValidateRequest  = "failed to validate request"
	ErrMsgFailedToAuthenticateUser = "failed to authenticate user"
	ErrMsgFailedToCreateEmergency  = "failed to create emergency"
	ErrMsgFailedToListEmergencies  = "failed to list emergencies"
	ErrMsgFailedToCreatePacient    = "failed to create pacient"
	ErrMsgFailedToFindPacient      = "failed to find pacient"
	ErrMsgFailedToCreateUser       = "failed to create user"
)
