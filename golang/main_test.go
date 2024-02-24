package main

import (
	"errors"
	"testing"
)

// Logger interface for logging messages
type Logger interface {
	Fatal(v ...interface{})
}

type MockLogger struct {
	fatalMessage string
}

func (m *MockLogger) Fatal(v ...interface{}) {
	m.fatalMessage = v[0].(string)
}

func TestHandleEnvLoadError(t *testing.T) {
	// Prepare a mock logger
	mockLogger := &MockLogger{}

	// Test case 1: Error is not nil
	err := errors.New("sample error")
	errReturned := HandleEnvLoadError(mockLogger, err) // Change the function name here
	if errReturned != err {
		t.Errorf("Expected error: %v, got: %v", err, errReturned)
	}

	// Test case 2: Error is nil
	err = nil
	errReturned = HandleEnvLoadError(mockLogger, err) // Change the function name here
	if errReturned != nil {
		t.Error("Expected nil error when error is nil")
	}
}

// HandleEnvLoadError is a function that handles errors when loading the .env file.
// It takes an error as a parameter and returns it if it's not nil.
func HandleEnvLoadError(logger Logger, err error) error { // Change the function name here
	if err != nil {
		return err
	}
	return nil
}
