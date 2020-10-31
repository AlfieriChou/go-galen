package routes

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

// Run 导出路由方法
func Run() {
	v1 := router.Group("/v1")
	addPingRoutes(v1)
	v2 := router.Group("/v2")
	addPingRoutes(v2)
	router.Run(":3000")
}
