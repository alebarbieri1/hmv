package errors

import (
	"errors"
	"fmt"
)

// Is returns true if two input errors are equal.
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// WithMessage formats a given error with a message. Output: "<message>: <err>"
func WithMessage(message string, err error) error {
	return fmt.Errorf("%s: %v", message, err)
}
