package logging

// Logger is an application logger that allows the logging of structured data
type Logger interface {
	// Debug logs at debug level
	Debug(message string, fields ...field)

	// Info logs at info level
	Info(message string, fields ...field)

	// Error logs at error level
	Error(message string, err error, fields ...field)
}

// Settings are logging settings that can be used to create customizable Loggers
type Settings struct {
	DevelopmentMode bool
}
