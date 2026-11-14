package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"example.com/issuepilot/internal/response"
)

// HealthCheck 处理服务健康检查请求
func HealthCheck(c *gin.Context) {
	response.Success(c, gin.H{
		"service": "issuepilot",
		"status":  "healthy",
	})
}

// ExampleError 示例：返回统一错误响应
func ExampleError(c *gin.Context) {
	response.Error(c, http.StatusBadRequest, "invalid request")
}
