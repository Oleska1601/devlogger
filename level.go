package devlogger

import "log/slog"

const (
	LevelFatal = slog.Level(12)
)

// Сопоставление уровней для форматирования в ReplaceAttr
var levelNames = map[slog.Level]string{
	slog.LevelDebug: "debug",
	slog.LevelInfo:  "info",
	slog.LevelWarn:  "warning",
	slog.LevelError: "error",
	LevelFatal:      "fatal",
}

var levelMap = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
	"fatal": LevelFatal,
}
