package logging

type MockLogger struct{}

func NewMockLogger() (Logger, error) {
	return new(MockLogger), nil
}

func (l *MockLogger) Debug(message string, fields ...field) {}

func (l *MockLogger) Info(message string, fields ...field) {}

func (l *MockLogger) Error(message string, err error, fields ...field) {}
