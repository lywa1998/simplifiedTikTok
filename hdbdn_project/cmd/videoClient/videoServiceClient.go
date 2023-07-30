package main

import (
	"fmt"
	"log"
	_"net"
	"context"
	"github.com/hdbdn77/simplifiedTikTok/pkg/videoService"
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
	publishActionServiceClient := videoService.NewPublishActionServiceClient(conn)
	publishListServiceClient := videoService.NewPublishListServiceClient(conn)

	// 3. 直接像调用本地方法一样调用GetProductStock方法
	publishActionResponse, err := publishActionServiceClient.PublishAction(context.Background(), &videoService.DouYinPublishActionRequest{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NCwidXNlcm5hbWUiOiJURVNUIiwiaXNzIjoidGVzdCIsImF1ZCI6WyJ6eXgiXSwiaWF0IjoxNjkwNzI4NjU0fQ.HgkYor2DHhVLZWfnlGE_5hJZl6kywXkXRXDMYUxy_M4", Data: []byte{},Title: "TEXT"})
	if err != nil {
		fmt.Println(err)
	}
	publishListResponse, _ := publishListServiceClient.PublishList(context.Background(), &videoService.DouYinPublishListRequest{ UserId: 4, Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6NCwidXNlcm5hbWUiOiJURVNUIiwiaXNzIjoidGVzdCIsImF1ZCI6WyJ6eXgiXSwiaWF0IjoxNjkwNzI4NjU0fQ.HgkYor2DHhVLZWfnlGE_5hJZl6kywXkXRXDMYUxy_M4"})
	

	// // fmt.Println("调用gRPC方法成功, StatusCode = ", resp.StatusCode)
	fmt.Println("投稿 :",publishActionResponse.StatusMsg)
	fmt.Println("查询 :",publishListResponse.StatusMsg)

	// 退出时关闭链接
	defer conn.Close()
}