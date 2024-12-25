package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化路由
	r := gin.Default()

	// TODO: 初始化数据库连接
	// TODO: 初始化定时任务
	// TODO: 初始化路由

	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Server startup failed: %v", err)
	}
}
