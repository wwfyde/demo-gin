package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func setRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return r
}

func main() {
	r := setRouter()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Run server failed: %v", err)
	}
}
