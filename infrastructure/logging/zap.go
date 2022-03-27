package logging

import (
	"fmt"

	"go.uber.org/zap"
)

type ZapLogger struct {
	core *zap.Logger
}

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

	return &ZapLogger{
		core: logger,
	}, nil
}

func (l *ZapLogger) Debug(message string, fields ...field) {
	l.core.Debug(message, toZapFields(fields)...)
}

func (l *ZapLogger) Info(message string, fields ...field) {
	l.core.Info(message, toZapFields(fields)...)
}

func (l *ZapLogger) Error(message string, err error, fields ...field) {
	l.core.Error(fmt.Sprintf("%s: %v", message, err), toZapFields(fields)...)
}

func toZapFields(fields []field) []zap.Field {
	zapFields := make([]zap.Field, 0, len(fields))

	for _, field := range fields {
		zapFields = append(zapFields, zap.String(field.Name, field.Value))
	}

	return zapFields
}
