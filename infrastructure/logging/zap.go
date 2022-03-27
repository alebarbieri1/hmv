package logging

import (
	"fmt"

	"go.uber.org/zap"
)

// ZapLogger is a Logger based on go.uber.org/zap logger
type ZapLogger struct {
	core *zap.Logger
}

// NewZapLogger creates a new ZapLogger
func NewZapLogger(settings *Settings) (Logger, error) {
	newLogger := zap.NewProduction

	if settings.DevelopmentMode {
		newLogger = zap.NewDevelopment
	}

	logger, err := newLogger(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	if err != nil {
		return nil, err
	}

	return &ZapLogger{core: logger}, nil
}

// Debug logs at debug level
func (l *ZapLogger) Debug(message string, fields ...field) {
	l.core.Debug(message, toZapFields(fields)...)
}

// Info logs at info level
func (l *ZapLogger) Info(message string, fields ...field) {
	l.core.Info(message, toZapFields(fields)...)
}

// Debug logs at error level
func (l *ZapLogger) Error(message string, err error, fields ...field) {
	l.core.Error(fmt.Sprintf("%s: %v", message, err), toZapFields(fields)...)
}

// toZapFields transforms a set of field into a set of zap.Field
func toZapFields(fields []field) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))

	for _, field := range fields {
		zapFields = append(zapFields, zap.String(field.Name, field.Value))
	}

	return zapFields
}
