package logger

import "log/slog"

func Debug(msg string, fields ...any) {
	slog.Debug(msg, fields...)
}

func Info(msg string, fields ...any) {
	slog.Info(msg, fields...)
}

func Warn(msg string, fields ...any) {
	slog.Warn(msg, fields...)
}

func Error(msg string, fields ...any) {
	slog.Error(msg, fields...)
}
