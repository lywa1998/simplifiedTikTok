package start_video_server

import (
	"fmt"
	"log"
	"net"

	"github.com/micro/simplifiedTikTok/apiserver/pkg/videoservice"
	"google.golang.org/grpc"
)

func Runserver() {
	server := grpc.NewServer()

	videoservice.RegisterPublishActionServiceServer(server, videoservice.PublishActionService)
	videoservice.RegisterPublishListServiceServer(server, videoservice.PublishListService)

	listener, err := net.Listen("tcp", ":8003")
	if err != nil {
		log.Fatalf("启动监听出错: %v", err)
	}

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("启动服务出错: %v", err)
	}

	fmt.Println("grpc服务器后台进行中...")

}