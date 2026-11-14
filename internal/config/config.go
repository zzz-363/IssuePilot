package config

import "os"

// GetPort 获取服务端口
func GetPort() string {
	port := os.Getenv("PORT")

	// 如果没有设置 PORT 环境变量，则使用默认端口
	if port == "" {
		port = "8080"
	}

	return port
}
