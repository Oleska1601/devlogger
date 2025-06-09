package devlogger

import (
	"log/slog"
)

type DebugLoggerMessage struct {
	Action  string      `json:"action"`
	Details interface{} `json:"details"`
}

func (l *Logger) Debug(action string, details interface{}) {
	message := DebugLoggerMessage{
		Action:  action,
		Details: details,
	}
	slog.Debug("", slog.Any("message", message))
}
