package devlogger

import (
	"log/slog"
	"strings"
)

type Logger struct {
	logger *slog.Logger
}

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
