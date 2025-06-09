package devlogger

import "log/slog"

type ErrorLoggerMessage struct {
	Action  string      `json:"action"`
	Error   error       `json:"error"`
	Details interface{} `json:"details"`
}

func (l *Logger) Error(action string, err error, details interface{}) {
	message := ErrorLoggerMessage{
		Action:  action,
		Error:   err,
		Details: details,
	}
	l.logger.Error("", slog.Any("message", message))
}
