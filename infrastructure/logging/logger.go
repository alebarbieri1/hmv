package logging

import "strconv"

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
