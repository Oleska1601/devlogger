package main

import (
	"devlogger/devlogger"
)

func main() {
	techLogger := devlogger.New("warning")

	techLogger.Debug("user authentication", map[string]interface{}{
		"user_id": 123,
		"method":  "oauth",
	})

	techLogger.Warning("user authentication", "expe", "act", "rea")
}
