package entity

import (
	"errors"
	"flavioltonon/hmv/domain/valueobject"
	"fmt"
)

var ErrNotFound = errors.New("not found")

func ErrProfileAlreadySet(profile valueobject.Profile) error {
	return fmt.Errorf("profile %s already set", profile)
}
