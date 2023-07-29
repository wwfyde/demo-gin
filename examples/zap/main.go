package main

import (
	"go.uber.org/zap"
	"time"
)

func main() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	for {
		logger.Info("My First Structured log with zap!",
			zap.String("hello", "世界!"),
			zap.Int("attempt", 3),
			zap.Duration("backoff", time.Second),
		)
		time.Sleep(time.Second)
	}
}
