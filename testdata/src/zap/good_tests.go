package zap

import (
	"go.uber.org/zap"
)

func GoodExamples() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	logger.Info("starting server on port 8080")
	logger.Error("failed to connect to database")
	logger.Info("user authenticated", zap.String("user", "alice"))

	sugar.Warn("low disk space")
	sugar.Debug("processing request")
}
