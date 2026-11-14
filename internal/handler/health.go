package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HealthCheck 处理服务健康检查请求
func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "ok",
		"data": gin.H{
			"service": "issuepilot",
			"status":  "healthy",
		},
	})
}
