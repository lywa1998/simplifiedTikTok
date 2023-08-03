package main

import (
	"context"
	"fmt"
	_ "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hdbdn77/simplifiedTikTok/pkg/startserver"

	"github.com/hdbdn77/simplifiedTikTok/pkg/clientconnect"
	"github.com/hdbdn77/simplifiedTikTok/pkg/userService"
)


type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	registerServiceClient := <- clientconnect.UserRegisterChan
	registerResponse, err := registerServiceClient.Register(context.Background(), &userService.DouYinUserRegisterRequest{Username: username, Password: password})
	clientconnect.UserRegisterChan <- registerServiceClient
	if err != nil {
		fmt.Println(err)
	}
	if registerResponse == nil {
		c.JSON(200, gin.H{
			"message": "err",
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{
			StatusCode: registerResponse.StatusCode, 
			StatusMsg: registerResponse.StatusMsg,
		},
		UserId:   registerResponse.UserId,
		Token:    registerResponse.Token,
	})

}

func init() {
	go startserver.StartUserServer()
	go startserver.StartVideoServer()
	go startserver.StartFeedServer()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/user/register", Register)
	r.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
