package test

import (
	log "logger"
)

func TestLogs() {
	password := "password123"
	apiKey := "secret123"
	token := "token123"

	log.Info("user password: " + password) // want "log message contains potentially sensitive variable: password" "log messages must be in English and contain no special characters or emojis"
	log.Debug("api_key=" + apiKey)         // want "log message contains potentially sensitive variable: apiKey" "log messages must be in English and contain no special characters or emojis"
	log.Info("token: " + token)            // want "log message contains potentially sensitive variable: token" "log messages must be in English and contain no special characters or emojis"
	log.Info("user authenticated successfully")
	log.Debug("api request completed")
	log.Info("token validated")

}
