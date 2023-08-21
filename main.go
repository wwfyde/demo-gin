package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/wwfyde/demo-gin/docs"
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
	v1 := r.Group("v1")
	v1.GET("/")

	return r
}

func init() {
	// TODO init
}

func main() {
	// 强制日志颜色化
	gin.ForceConsoleColor()
	//// 禁止日志的颜色
	//gin.DisableConsoleColor()

	// Parse Command line flags
	flag.Parse()
	//docs.SwaggerInfo.BasePath = "api/v1"
	r := setRouter()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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

	// serve

}
