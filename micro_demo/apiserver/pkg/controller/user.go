package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/simplifiedTikTok/apiserver/pkg/clientconnect"
	"github.com/micro/simplifiedTikTok/apiserver/pkg/userservice"
	// "github.com/micro/simplifiedTikTok/userservice/pkg/userservice"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	registerServiceClient := <- clientconnect.UserRegisterChan
	registerResponse, err := registerServiceClient.Register(context.Background(), &userservice.DouYinUserRegisterRequest{Username: username, Password: password})
	clientconnect.UserRegisterChan <- registerServiceClient
	if err != nil {
		fmt.Println(err)
	}
	if registerResponse == nil {
		c.JSON(http.StatusOK, gin.H{
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