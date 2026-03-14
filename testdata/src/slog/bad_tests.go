package slog

import (
	"log/slog"
)

var (
	password = "password"
	apiKey   = "qwerty"
	token    = "12345"
)

func BadExamples() {
	slog.Info("Starting server on port 8080") // want "log should start with a lowercase letter"
	slog.Error("запуск сервера")              // want "log should contain only English letters"
	slog.Warn("server started! 🚀")            // want "log should not contain emoji"
	slog.Info("user password: " + password)   // want "log may contain sensitive data"

	logger := slog.New(nil)
	logger.Info("Starting server") // want "log should start with a lowercase letter"
	logger.Debug("token=" + token) // want "log may contain sensitive data"
	logger.Info("ok!!")            // want "log contains disallowed special characters"
}
