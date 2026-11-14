package main

import (
	"log"

	"example.com/issuepilot/internal/config"
	"example.com/issuepilot/internal/router"
)

func main() {
	// 获取服务端口
	port := config.GetPort()

	// 初始化路由
	r := router.SetupRouter()

	// 启动 HTTP 服务
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
