package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micro/simplifiedTikTok/apiserver/pkg/route"

	"github.com/micro/simplifiedTikTok/userservice/cmd/start_user_server"
	"github.com/micro/simplifiedTikTok/videoservice/cmd/start_video_server"
	"github.com/micro/simplifiedTikTok/feedservice/cmd/start_feed_server"

)

func init(){
	go start_user_server.Runserver()
	go start_video_server.Runserver()
	go start_feed_server.Runserver()
}

func main() {
	r := gin.Default()

	route.InitRouter(r)
	
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
