package entity

import "errors"

var (
	ErrNotFound                  = errors.New("not found")
	ErrUserAlreadyIsAPacient     = errors.New("user already is a pacient")
	ErrUsernameAlreadyInUse      = errors.New("username is already in use")
	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
)
