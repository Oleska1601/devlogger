package devlogger

import (
	"context"
	"log/slog"
	"os"
)

type FatalLoggerMessage struct {
	Action  string      `json:"action"`
	Error   error       `json:"error"`
	Details interface{} `json:"details"`
}

func (l *Logger) Fatal(action string, err error, details interface{}) {
	message := FatalLoggerMessage{
		Action:  action,
		Error:   err,
		Details: details,
	}
	l.logger.Log(context.Background(), LevelFatal, "", slog.Any("message", message))
	os.Exit(1)
}
