package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/wwfyde/demo-gin/api"
	v12 "github.com/wwfyde/demo-gin/api/v1"
	_ "github.com/wwfyde/demo-gin/docs"
	"go.uber.org/zap"
	"log"
	"net/http"
	"time"
)

var (
	port = flag.Int("port", 8000, "Server listening port")
)

func init() {
	// TODO init
}

// @title Demo Gin API
// @version 3.1.0
// @host localhost:8000
// @BasePath /api/v1
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	// 强制日志颜色化
	gin.ForceConsoleColor()
	//// 禁止日志的颜色
	//gin.DisableConsoleColor()

	// Parse Command line flags
	flag.Parse()
	//docs.SwaggerInfo.BasePath = "/api/v1"
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		hello := v1.Group("/hello")
		{
			hello.GET("", v12.HelloWorld)
		}
	}
	v1.GET("/ping", api.Ping)
	v1.GET("", func(c *gin.Context) { c.String(http.StatusOK, "api/v1") })
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
