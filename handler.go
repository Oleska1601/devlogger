package devlogger

import (
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
)

// newHandler creates a customized slog.Handler with:
//   - JSON output format
//   - Custom log level names
//   - Formatted timestamps
//   - Simplified source file paths
func newHandler(level slog.Level) slog.Handler {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Customize timestamp format
			if a.Key == slog.TimeKey {
				return slog.String(a.Key, a.Value.Time().Format("2006-01-02T15:04:05.000Z"))
			}
			// Customize level names
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := levelNames[level]
				if !exists {
					levelLabel = level.String()
				}
				return slog.String(a.Key, levelLabel)
			}
			// Customize source file format
			if a.Key == slog.SourceKey {
				source := a.Value.Any().(*slog.Source)
				return slog.String("file", filepath.Base(source.File)+":"+strconv.Itoa(source.Line))
			}
			// Remove current message key to set define message struct for all types of log level
			if a.Key == slog.MessageKey {
				return slog.Attr{}
			}

			return a
		},
	})
	return handler
}
