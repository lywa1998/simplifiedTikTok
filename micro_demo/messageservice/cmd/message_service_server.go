package main

import (
	"fmt"
	"log"
	"net"

	"github.com/micro/simplifiedTikTok/messageservice/pkg/messageservice"
	"google.golang.org/grpc"
)

func main() {
	server := grpc.NewServer()

	messageservice.RegisterMessageChatServiceServer(server, messageservice.MessageChatService)
	messageservice.RegisterMessageActionServiceServer(server, messageservice.MessageActionService)

	listener, err := net.Listen("tcp", ":8005")
	if err != nil {
		log.Fatalf("启动监听出错: %v", err)
	}

	err = server.Serve(listener)
	if err != nil {
		log.Fatalf("启动服务出错: %v", err)
	}

	fmt.Println("grpc服务器后台进行中...")

}
