package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func Logger(c *gin.Context) {
	log.Println("Before route execution: Logger")
	c.Next()
	log.Println("After route execution: Logger")
}

func Authorize(c *gin.Context) {
	// dummy authorization logic
	isAuthorized := true

	if !isAuthorized {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		c.Abort() // it's important to abort the request here to prevent it from going further
		return
	}
	c.Next()
}

func main() {
	r := gin.Default()

	// global use middleware
	r.Use(Logger)

	r.GET("/public", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "This is a public endpoint"})
	})

	authorized := r.Group("/secret")

	{
		authorized.Use(Authorize)
		authorized.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Welcome to secret page"})
		})
	}

	r.Run()
}
