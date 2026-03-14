package zap

import (
	"go.uber.org/zap"
)

var (
	password = "password"
	apiKey   = "abcabc"
	token    = "qwertytest"
)

func BadExamples() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	logger.Info("Starting server on port 8080")       // want "log should start with a lowercase letter"
	logger.Error("ошибка подключения")                // want "log should contain only English letters"
	logger.Info("server started! 🚀")                  // want "log should not contain emoji"
	logger.Info("api_key=" + apiKey)                  // want "log may contain sensitive data"
	logger.Error("Failed to connect", zap.Error(nil)) // want "log should start with a lowercase letter"

	sugar.Info("token: " + token)                  // want "log may contain sensitive data"
	sugar.Info("ssn: 123-45-6789")                 // want "log may contain sensitive data"
	sugar.Warn("warning: something went wrong...") // want "log contains disallowed special characters"
	sugar.Info("ok!!!")                            // want "log contains disallowed special characters"
	sugar.Debug("отладочное сообщение")            // want "log should contain only English characters"
	sugar.Error("Debug message")                   // want "log should start with a lowercase letter"

}
