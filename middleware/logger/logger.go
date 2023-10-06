package logger

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {

		t := time.Now()
		// TODO
		//time.Sleep(3 * time.Millisecond)
		log, _ := zap.NewProduction()
		// before request
		c.Next()
		// after request
		latency := time.Since(t)
		log.Info("latency",
			zap.String("latency", fmt.Sprintf("%13v", latency)),
		)

	}
}
