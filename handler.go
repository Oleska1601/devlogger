package devlogger

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"
)

func getFileName(file string, line int) string {
	fileName := filepath.Base(file)
	// Разделяем строку и находим 'devlogger'
	parts := strings.Split(file, "/devlogger/")
	if len(parts) > 1 {
		// Возвращаем часть с 'devlogger' и следующим путем
		filePath := parts[1]
		return fmt.Sprintf("%s:%d", filePath, line)
	} else {
		return fmt.Sprintf("%s:%d", fileName, line)
	}
}

func newHandler(level slog.Level) slog.Handler {
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     level,
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			// Кастомизация формата времени
			if a.Key == slog.TimeKey {
				return slog.String(a.Key, a.Value.Time().Format("2006-01-02T15:04:05.000Z"))
			}

			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := levelNames[level]
				if !exists {
					levelLabel = level.String()
				}
				return slog.String(a.Key, levelLabel)
			}
			// Кастомизация формата источника
			if a.Key == slog.SourceKey {
				source := a.Value.Any().(*slog.Source)
				filename := getFileName(source.File, source.Line)
				return slog.String("file", filename)
			}

			if a.Key == slog.MessageKey {
				return slog.Attr{}
			}

			return a
		},
	})
	return handler
}
