package devlogger

import "log/slog"

// LevelFatal defines the custom log level for logging a fatal level.
const (
	LevelFatal = slog.Level(12)
)

// levelNames provides mapping between slog.Level values and their string representations
var levelNames = map[slog.Level]string{
	slog.LevelDebug: "debug",
	slog.LevelInfo:  "info",
	slog.LevelWarn:  "warning",
	slog.LevelError: "error",
	LevelFatal:      "fatal",
}

// levelMap provides mapping between string level names and slog.Level values
var levelMap = map[string]slog.Level{
	"debug": slog.LevelDebug,
	"info":  slog.LevelInfo,
	"warn":  slog.LevelWarn,
	"error": slog.LevelError,
	"fatal": LevelFatal,
}

// DebugLoggerMessage represents the structure of message for logging in debug level.
type DebugLoggerMessage struct {
	Action  string      `json:"action"`  // Action shows the action that is being logged.
	Details interface{} `json:"details"` // Details shows the additional debug details.
}

// LoggerMessageInfoStatus defines possible status values for logging in info level.
type LoggerMessageInfoStatus string

// Constants for LoggerMessageInfoStatus values
const (
	InfoStatusTry     LoggerMessageInfoStatus = "try"     // InfoStatusTry shows that action is in progress
	InfoStatusSuccess LoggerMessageInfoStatus = "info"    // InfoStatusSuccess shows that action completed successfully
	InfoStatusFailure LoggerMessageInfoStatus = "failure" // InfoStatusFailure shows that action failed
)

// InfoLoggerMessage represents the structure of message for logging in info level.
type InfoLoggerMessage struct {
	Action  string                  `json:"action"`  // Action shows the action that is being logged.
	Status  LoggerMessageInfoStatus `json:"status"`  // Status shows the status of the action.
	Details interface{}             `json:"details"` // Details shows the additional info details.
}

// LoggerMessageWarningDetails represents the structure of message details for logging in warning level.
type LoggerMessageWarningDetails struct {
	Expected string `json:"expected"` // Expected shows the expected behavior of the program.
	Actual   string `json:"actual"`   // Actual shows the real behavior of the program.
	Reason   string `json:"reason"`   // Reason shows the possible reason of not matching expected and actual behaviors of the program.
}

// LoggerMessageWarningDetails represents the structure of message for logging in warning level.
type WarningLoggerMessage struct {
	Action  string                      `json:"action"`  // Action shows the action that is being logged.
	Details LoggerMessageWarningDetails `json:"details"` // Details shows the additional warning details.
}

// ErrorLoggerMessage represents the structure of message for logging in error level.
type ErrorLoggerMessage struct {
	Action  string      `json:"action"`  // Action shows the action that is being logged.
	Error   string      `json:"error"`   // Error shows the error of the action.
	Details interface{} `json:"details"` // Details shows the additional error details.
}

// FatalLoggerMessage represents the structure of message for logging in fatal level.
type FatalLoggerMessage struct {
	Action  string      `json:"action"`  // Action shows the action that is being logged.
	Error   string      `json:"error"`   // Error shows the error of the action.
	Details interface{} `json:"details"` // Details shows the additional fatal details.
}
