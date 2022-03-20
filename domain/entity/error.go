package entity

import (
	"errors"
	"flavioltonon/hmv/domain/valueobject"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func ErrProfileKindAlreadySet(profileKind valueobject.ProfileKind) error {
	return fmt.Errorf("profile kind already set as %s", profileKind)
}

func ErrInvalidStatusChange(from, to valueobject.EmergencyStatus) error {
	return fmt.Errorf("status cannot be changed from %s to %s", from, to)
}
