package router

import (
	"github.com/gin-gonic/gin"

	"example.com/issuepilot/internal/handler"
)

// SetupRouter 初始化并配置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 健康检查接口
	r.GET("/api/v1/health", handler.HealthCheck)

	return r
}
