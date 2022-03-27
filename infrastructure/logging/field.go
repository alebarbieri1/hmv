package logging

import (
	"strconv"
)

// field defines a key/value pair that can be used with Logger logging
type field struct {
	Name  string
	Value string
}

// String creates a new field with a typed string value
func String(name, value string) field {
	return field{
		Name:  name,
		Value: value,
	}
}

// Int creates a new field with a typed int value
func Int(name string, value int) field {
	return field{
		Name:  name,
		Value: strconv.Itoa(value),
	}
}

// Error creates a new field with a typed error value and a default key "error"
func Error(value error) field {
	return field{
		Name:  "error",
		Value: value.Error(),
	}
}
