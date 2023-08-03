package main

import (
	"fmt"
	"log"
	_"net"
	"time"
	"context"
	"github.com/hdbdn77/simplifiedTikTok/pkg/feedService"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 1. 新建连接，端口是服务端开放的8003端口
	// 没有证书会报错
	conn, err := grpc.Dial(":8004", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	feedServiceClient := feedService.NewFeedServiceClient(conn)

	// 3. 直接像调用本地方法一样调用GetProductStock方法
	time := time.Now().Unix()
	fmt.Println(time)
	feedResponse, err := feedServiceClient.Feed(context.Background(), &feedService.DouYinFeedRequest{Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6MSwidXNlcm5hbWUiOiJURVNUIiwiaXNzIjoidGVzdCIsImF1ZCI6WyJ6eXgiXSwiaWF0IjoxNjkwOTY2NzExfQ.bYe7k3nfZ-TpEOD7TEH3fBEKP4A-bYd5dckpgDDg37k", LatestTime: time})
	if err != nil {
		fmt.Println(err)
	}

	// // fmt.Println("调用gRPC方法成功, StatusCode = ", resp.StatusCode)
	fmt.Println("查询 :",feedResponse.StatusMsg)
	fmt.Println(feedResponse.VideoList)

	// 退出时关闭链接
	defer conn.Close()
}