package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/micro/simplifiedTikTok/apiserver/pkg/clientconnect"
	"github.com/micro/simplifiedTikTok/apiserver/pkg/userservice"
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
			"message": "register err",
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

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	loginServiceClient := <- clientconnect.UserLoginChan
	loginResponse, err := loginServiceClient.Login(context.Background(), &userservice.DouYinUserLoginRequest{Username: username, Password: password})
	clientconnect.UserLoginChan <- loginServiceClient
	if err != nil {
		fmt.Println(err)
	}
	if loginResponse == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "login err",
		})
		return
	}
	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{
			StatusCode: loginResponse.StatusCode,
			StatusMsg: loginResponse.StatusMsg,
		},
		UserId:   loginResponse.UserId,
		Token:    loginResponse.Token,
	})
}

func UserInfo(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	token := c.Query("token")
	userServiceClient := <- clientconnect.UserChan
	userResponse, err := userServiceClient.Find(context.Background(), &userservice.DouYinUserRequest{UserId: userId, Token: token})
	clientconnect.UserChan <- userServiceClient
	if err != nil {
		fmt.Println(err)
	}
	if userResponse == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": "user info err",
		})
		return
	}
	c.JSON(http.StatusOK, UserResponse{
		Response: Response{
			StatusCode: userResponse.StatusCode,
			StatusMsg: userResponse.StatusMsg,
		},
		User: User{
			Id: userResponse.User.Id,
			Name: userResponse.User.Name,
			FollowCount: userResponse.User.FollowCount,
			FollowerCount: userResponse.User.FollowerCount,
			IsFollow: userResponse.User.IsFollow,
			Avatar: userResponse.User.Avatar,
			BackgroundImage: userResponse.User.BackgroundImage,
			Signature: userResponse.User.Signature,
			TotalFavorited: userResponse.User.TotalFavorited,
			WorkCount: userResponse.User.WorkCount,
			FavoriteCount: userResponse.User.FavoriteCount,
		},
	})
}