package test

import (
	log "logger"
)

func TestEmojiAndSymbolsLogs() {
	log.Info("server started!🚀")                 // want "log messages must be in English and contain no special characters or emojis"
	log.Error("connection failed!!!")            // want "log messages must be in English and contain no special characters or emojis"
	log.Warn("warning: something went wrong...") // want "log messages must be in English and contain no special characters or emojis"

	log.Info("server started")
	log.Error("connection failed")
	log.Warn("something went wrong")
}
