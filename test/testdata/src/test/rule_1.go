package test

import (
	"log/slog"

	log "logger"
)

func TestFistLetterLogs() {
	log.Info("Starting server on port 8080")    // want "log messages must start with a lowercase letter"
	slog.Error("Failed to connect to database") // want "log messages must start with a lowercase letter"

	log.Info("starting server on port 8080")
	slog.Error("failed to connect to database")
}
