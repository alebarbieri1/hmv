package logging

import (
	"fmt"
	"strconv"
)

type Logger interface {
	Debug(message string, fields ...field)
	Info(message string, fields ...field)
	Error(message string, err error, fields ...field)
}

type field struct {
	Name  string
	Value interface{}
}

func String(name, value string) field {
	return field{
		Name:  name,
		Value: value,
	}
}

func Stringer(name string, value fmt.Stringer) field {
	return field{
		Name:  name,
		Value: value.String(),
	}
}

func Int(name string, value int) field {
	return field{
		Name:  name,
		Value: strconv.Itoa(value),
	}
}

func Error(value error) field {
	return field{
		Name:  "error",
		Value: value.Error(),
	}
}

type Settings struct {
	DevelopmentMode bool
}
