package logging

// NopLogger is a Logger that performs no actions.
type NopLogger struct{}

// NewNopLogger creates a new NopLogger
func NewNopLogger() Logger { return new(NopLogger) }

// Debug logs at debug level
func (l *NopLogger) Debug(message string, fields ...field) {}

// Info logs at info level
func (l *NopLogger) Info(message string, fields ...field) {}

// Error logs at error level
func (l *NopLogger) Error(message string, err error, fields ...field) {}
