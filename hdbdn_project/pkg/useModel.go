package main

import (
	_ "crypto/rand"
	_ "crypto/rsa"
	"fmt"
	_ "github.com/golang-jwt/jwt/v5"
	"github.com/hdbdn77/simplifiedTikTok/pkg/model"
	_"github.com/hdbdn77/simplifiedTikTok/pkg/utils"
)

func main() {

	// // user功能测试
	// user := model.User{Username: "LPF", Password: "ZYX"}
	// userA, err := model.Register(&user)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(*userA)
	// }
	// user = model.User{Username: "LPF", Password: "zyx"}
	// userB, _ := model.FindUserByUsername(&model.User{Username: user.Username})
	// if userB.Password == user.Password {
	// 	fmt.Println(*userB)
	// } else {
	// 	fmt.Println("Wrong password")
	// }

	// // token功能测试
	// token, err := utils.GenToken(userB.Id, userB.Username)
	// if err != nil {
	// 	fmt.Println("生成出错误")
	// }
	// fmt.Println(token)

	// claims, err := utils.ParseToken(token)
	// if err != nil {
	// 	fmt.Println("生成出错误")
	// }
	// fmt.Println(claims.ID)
	// fmt.Println(claims.Username)

	// video功能测试
	video := model.Video{AuthorId: 4, PlayUrl: "https://www.baidu.com", CoverUrl: "https://www.baidu.com", Title: "test"}
	_, err := model.CreateVideo(&video)
	if err != nil {
		fmt.Println(err)
	}
	user := model.User{Id: 4}
	_, err = model.AddWorkCount(&user)
	if err != nil {
		fmt.Println(err)
	}


	// video = model.Video{AuthorId: 4, PlayUrl: "https://www.lpf.com", CoverUrl: "https://www.zyx.com", Title: "Test"}
	// err = model.CreateVideo(&video)
	// if err != nil {
	// 	fmt.Println(err)
	// } 
	// videolist, err := model.ListVideo(4)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(*videolist)

}
