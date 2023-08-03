package main

import (
	"fmt"
	"log"
	_"net"
	"context"
	"github.com/micro/simplifiedTikTok/videoservice/pkg/videoservice"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. 新建连接，端口是服务端开放的8003端口
	// 没有证书会报错
	conn, err := grpc.Dial(":8003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	publishActionServiceClient := videoservice.NewPublishActionServiceClient(conn)
	publishListServiceClient := videoservice.NewPublishListServiceClient(conn)

	// 3. 直接像调用本地方法一样调用GetProductStock方法
	publishActionResponse1, err := publishActionServiceClient.PublishAction(context.Background(), &videoservice.DouYinPublishActionRequest{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJURVNUIiwiaXNzIjoidGVzdCIsImF1ZCI6WyJ6eXgiXSwiaWF0IjoxNjkwOTY2NzExfQ.bYe7k3nfZ-TpEOD7TEH3fBEKP4A-bYd5dckpgDDg37k", Data: []byte{},Title: "TEXT4"})

	if err != nil {
		fmt.Println(err)
	}
	publishListResponse, _ := publishListServiceClient.PublishList(context.Background(), &videoservice.DouYinPublishListRequest{ UserId: 1, Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJURVNUIiwiaXNzIjoidGVzdCIsImF1ZCI6WyJ6eXgiXSwiaWF0IjoxNjkwOTY2NzExfQ.bYe7k3nfZ-TpEOD7TEH3fBEKP4A-bYd5dckpgDDg37k"})
	

	// // fmt.Println("调用gRPC方法成功, StatusCode = ", resp.StatusCode)
	fmt.Println("投稿 :",publishActionResponse1.StatusMsg)
	// fmt.Println("投稿 :",publishActionResponse2.StatusMsg)
	// fmt.Println("投稿 :",publishActionResponse3.StatusMsg)
	fmt.Println("查询 :",publishListResponse.StatusMsg)
	fmt.Println(publishListResponse.VideoList[0].Author.FollowCount)

	// 退出时关闭链接
	defer conn.Close()
}