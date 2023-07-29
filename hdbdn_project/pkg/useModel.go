package main

import (
	"github.com/hdbdn77/simplifiedTikTok/pkg/model"
	"github.com/hdbdn77/simplifiedTikTok/pkg/utils"
	"fmt"
	_"github.com/golang-jwt/jwt/v5"
	_"crypto/rsa"
	_"crypto/rand"
)

func main() {
	user := model.User{Username: "LPF", Password: "XYZ"}
	userA, err := model.Register(&user)
	if err != nil {
		fmt.Println(err)
	}else{
		fmt.Println(*userA)
	}
	user = model.User{Username: "LPF", Password: "zyx"}
	userB, _ := model.Login(&model.User{Username: user.Username})
	if userB.Password == user.Password {
		fmt.Println(*userB)
	}else{
		fmt.Println("Wrong password")
	}

	token, err := utils.GenToken(userB.Id, userB.Username)
	if err != nil {
		fmt.Println("生成出错误")
	}
	fmt.Println(token)

	claims, err := utils.ParseToken(token)
	if err != nil {
		fmt.Println("生成出错误")
	}
	fmt.Println(claims.ID)
	fmt.Println(claims.Username)

}