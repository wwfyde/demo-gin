package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

var (
	port = flag.Int("port", 8000, "Server listening port")
)

func setRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func main() {
	// 强制日志颜色化
	gin.ForceConsoleColor()
	//// 禁止日志的颜色
	//gin.DisableConsoleColor()
	flag.Parse()
	r := setRouter()

	r.Use()
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	logger.Info("My First Structured log with zap!",
		zap.String("hello", "世界!"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
	if err := r.Run(fmt.Sprintf(":%d", *port)); err != nil {
		log.Fatalf("Run server failed: %v", err)
	}
}
