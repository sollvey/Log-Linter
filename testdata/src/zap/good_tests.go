package zap

import (
	"go.uber.org/zap"
)

func GoodExamples() {
	logger, _ := zap.NewProduction()
	sugar := logger.Sugar()

	logger.Info("starting server on port 8080")
	logger.Error("failed to connect to database")

	sugar.Warn("low disk space")
	sugar.Debug("processing request")

	// С полями
	logger.Info("user authenticated", zap.String("user", "alice"))
	logger.Error("connection timeout", zap.Duration("timeout", 5))
	sugar.Infow("cache miss", "key", "user:123")
	sugar.Debugw("api call completed", "status", 200)
}
