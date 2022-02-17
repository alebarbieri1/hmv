package errors

import (
	"errors"
	"fmt"
)

func Is(err, target error) bool {
	return errors.Is(err, target)
}

func WithMessage(message string, err error) error {
	return fmt.Errorf("%s: %v", message, err)
}
