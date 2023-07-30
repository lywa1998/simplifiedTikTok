package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hdbdn77/simplifiedTikTok/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()

	service.RegisterRegisterServiceServer(server, service.RegisterService)
	service.RegisterLoginServiceServer(server, service.LoginService)
	service.RegisterUserServiceServer(server, service.UserService)

	listener, err := net.Listen("tcp", ":8002")
	if err != nil {
		log.Fatalf("启动监听出错: %v", err)
	}

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("启动服务出错: %v", err)
	}

	fmt.Println("grpc服务器后台进行中...")

}
