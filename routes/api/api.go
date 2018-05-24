package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Routes(route *gin.Engine) {
	api := route.Group("/api")
	api.GET("/user.json", func(c *gin.Context) {
		c.JSON(http.StatusOK, []gin.H{
			gin.H{"name": "Name A"},
			gin.H{"name": "Name B"},
			gin.H{"name": "Name C"},
		})
	})
}
