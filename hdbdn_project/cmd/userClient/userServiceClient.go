package main

import (
	"fmt"
	"log"
	_"net"
	"context"
	"github.com/hdbdn77/simplifiedTikTok/pkg/userService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. 新建连接，端口是服务端开放的8002端口
	// 没有证书会报错
	conn, err := grpc.Dial(":8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	registerServiceClient := userService.NewRegisterServiceClient(conn)
	loginServiceClient := userService.NewLoginServiceClient(conn)
	userServiceClient := userService.NewUserServiceClient(conn)

	// 3. 直接像调用本地方法一样调用GetProductStock方法
	registerResponse, _ := registerServiceClient.Register(context.Background(), &userService.DouYinUserRegisterRequest{Username: "TEST", Password: "TEST"})
	loginResponse, _ := loginServiceClient.Login(context.Background(), &userService.DouYinUserLoginRequest{Username: "TEST", Password: "TEST"})
	fmt.Println("token :",loginResponse.Token)
	userResponse, _ := userServiceClient.Find(context.Background(), &userService.DouYinUserRequest{UserId: 4, Token: loginResponse.Token})

	// fmt.Println("调用gRPC方法成功, StatusCode = ", resp.StatusCode)
	fmt.Println("注册 :",registerResponse.StatusMsg)
	fmt.Println("登录 :",loginResponse.StatusMsg)
	fmt.Println("查询 :",userResponse.StatusMsg)

	// 退出时关闭链接
	defer conn.Close()
}