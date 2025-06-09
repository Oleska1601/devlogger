package devlogger

import "log/slog"

// for info
type LoggerMessageInfoStatus string

const (
	InfoStatusTry     LoggerMessageInfoStatus = "try"
	InfoStatusSuccess LoggerMessageInfoStatus = "info"
	InfoStatusFailure LoggerMessageInfoStatus = "failure"
)

type InfoLoggerMessage struct {
	Action  string                  `json:"action"`
	Status  LoggerMessageInfoStatus `json:"status"`
	Details interface{}             `json:"details"`
}

func (l *Logger) Info(action string, status LoggerMessageInfoStatus, details interface{}) {
	message := InfoLoggerMessage{
		Action:  action,
		Status:  status,
		Details: details,
	}

	l.logger.Info("", slog.Any("message", message))
}
