package routes

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// CreateRouter 创建路由方法
func CreateRouter() *gin.Engine {
	v1 := router.Group("/v1")
	addPingRoutes(v1)
	v2 := router.Group("/v2")
	addPingRoutes(v2)
	return router
}
