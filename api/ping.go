package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Ping print ping
//
//	@Summary	ping and pong
//	@Produce	json
//	@Success	200
//	@Router		/ping [get]
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]any{"msg": "pong"})
}
