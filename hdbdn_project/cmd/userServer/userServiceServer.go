package main

import (
	"fmt"
	"log"
	"net"

	"github.com/hdbdn77/simplifiedTikTok/pkg/userService"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()

	userService.RegisterRegisterServiceServer(server, userService.RegisterService)
	userService.RegisterLoginServiceServer(server, userService.LoginService)
	userService.RegisterUserServiceServer(server, userService.UserService)

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
