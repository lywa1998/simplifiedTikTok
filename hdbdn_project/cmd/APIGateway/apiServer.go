package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hdbdn77/simplifiedTikTok/pkg/startServer"
)


func init() {
	go startServer.StartUserServer()
	go startServer.StartVideoServer()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}