package devlogger

import "log/slog"

type LoggerMessageWarningDetails struct {
	Expected string `json:"expected"`
	Actual   string `json:"actual"`
	Reason   string `json:"reason"`
}

type WarningLoggerMessage struct {
	Action  string                      `json:"action"`
	Details LoggerMessageWarningDetails `json:"details"`
}

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
