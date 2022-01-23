package logging

type Logger interface {
	Debug(message string, fields ...field)
	Info(message string, fields ...field)
	Error(message string, err error, fields ...field)
}

type field struct {
	Name  string
	Value interface{}
}

func Field(name, value string) field {
	return field{
		Name:  name,
		Value: value,
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
