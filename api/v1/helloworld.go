package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HelloWorld godoc
//
// @Summary 	hello world
// @Tags		example
// @Accept		json
// @Success		200
// @Router		/hello [get]
func HelloWorld(c *gin.Context) {
	v := "世界"
	c.JSON(http.StatusOK, map[string]any{"Hello": v, "hello": "world"})
}
