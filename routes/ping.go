package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Data 返回data结构体
type Data struct {
	message string
}

func addPingRoutes(rg *gin.RouterGroup) {
	ping := rg.Group("/ping")

	ping.GET("/", func(c *gin.Context) {
		var data Data
		data.message = "pong"
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "ok",
			"data":    data,
		})
	})
}
