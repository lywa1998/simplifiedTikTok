package startServer

import (
	"fmt"
	"log"
	"net"

	"github.com/hdbdn77/simplifiedTikTok/pkg/videoService"
	"google.golang.org/grpc"
)

func StartVideoServer() {
	server := grpc.NewServer()

	videoService.RegisterPublishActionServiceServer(server, videoService.PublishActionService)
	videoService.RegisterPublishListServiceServer(server, videoService.PublishListService)

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