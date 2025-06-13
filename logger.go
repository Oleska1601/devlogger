// Package devlogger provides structured logger with multiple log levels.
// It wraps the standard slog package with additional functionality.
//
// This package supports the following levels:
//   - Debug
//   - Info
//   - Warning
//   - Error
//   - Fatal
//
// Each log level has its own structured message format.
package devlogger

import (
	"context"
	"log/slog"
	"os"
	"strings"
)

// LoggerInterface defines the methods for logger.
type LoggerInterface interface {
	// Debug logs a debug level message with action and details.
	Debug(action string, details interface{})
	// Info logs an info level message with action, status and details.
	Info(action string, status LoggerMessageInfoStatus, details interface{})
	// Warning logs a warning level message with action, expected, actual values and reason.
	Warning(action string, expected string, actual string, reason string)
	// Error logs an error level message with action, error and details.
	Error(action string, err error, details interface{})
	// Fatal logs a fatal level message with action, error and details and terminates the program.
	Fatal(action string, err error, details interface{})
}

// Logger is the concrete implementation of LoggerInterface.
type Logger struct {
	logger *slog.Logger // logger is a pointer to base slog.Logger
}

// New creates a new Logger with the specified log level.
//
// The inputLevel parameter accepts string values: "debug", "info", "warning", "error", "fatal".
// If an unknown level is provided, it defaults to "info".
func New(inputLevel string) *Logger {
	level, ok := levelMap[strings.ToLower(inputLevel)]
	if !ok {
		level = slog.LevelInfo
	}
	handler := newHandler(level)
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return &Logger{
		logger: logger,
	}
}

// Debug logs a debug level message with action and details.
//
// The details parameter can be any structured data that will be included in the log.
func (l *Logger) Debug(action string, details interface{}) {
	message := DebugLoggerMessage{
		Action:  action,
		Details: details,
	}
	slog.Debug("", slog.Any("message", message))
}

// Info logs an info level message with action, status and details.
//
// The status parameter should be one of the predefined LoggerMessageInfoStatus values.
// The details parameter can be any structured data that will be included in the log.
func (l *Logger) Info(action string, status LoggerMessageInfoStatus, details interface{}) {
	message := InfoLoggerMessage{
		Action:  action,
		Status:  status,
		Details: details,
	}

	l.logger.Info("", slog.Any("message", message))
}

// Warning logs a warning level message with action, expected, actual values and reason.
//
// This is typically used to log situations where something unexpected happened,
// but the app can continue running.
func (l *Logger) Warning(action string, expected, actual, reason string) {
	message := WarningLoggerMessage{
		Action: action,
		Details: LoggerMessageWarningDetails{
			Expected: expected,
			Actual:   actual,
			Reason:   reason,
		},
	}
	l.logger.Warn("", slog.Any("message", message))
}

// Error logs an error level message with action, error and details.
//
// The error parameter contains the error that occurred.
// The details parameter can be any structured data that will be included in the log.
func (l *Logger) Error(action string, err error, details interface{}) {
	message := ErrorLoggerMessage{
		Action:  action,
		Error:   err.Error(),
		Details: details,
	}
	l.logger.Error("", slog.Any("message", message))
}

// Fatal logs a fatal level message with action, error and details and terminates the program.
//
// The error parameter contains the error that occurred.
// The details parameter can be any structured data that will be included in the log.
func (l *Logger) Fatal(action string, err error, details interface{}) {
	message := FatalLoggerMessage{
		Action:  action,
		Error:   err.Error(),
		Details: details,
	}
	l.logger.Log(context.Background(), LevelFatal, "", slog.Any("message", message))
	os.Exit(1)
}
