package test

import (
	"log/slog"

	log "logger"
)

func TestEnglishLogs() {
	log.Info("запуск сервера")                     // want "log messages must be in English and contain no special characters or emojis"
	slog.Error("ошибка подключения к базе данных") // want "log messages must be in English and contain no special characters or emojis"

	log.Info("starting server")
	slog.Error("failed to connect to database")
}
