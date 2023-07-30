package videoService
import (
	"context"
	"fmt"
	"time"
	"os"

	"github.com/hdbdn77/simplifiedTikTok/pkg/model"
	"github.com/hdbdn77/simplifiedTikTok/pkg/utils"
)

var PublishActionService = &publishActionService{}
var PublishListService = &publishListService{}

type publishActionService struct {}
type publishListService struct {}

func (pA *publishActionService) PublishAction(context context.Context, request *DouYinPublishActionRequest) (*DouYinPublishActionResponse, error) {
	//实现具体的业务逻辑
	claims, _ := utils.ParseToken(request.Token)
	if (claims == nil) {
		return &DouYinPublishActionResponse{
			StatusCode: -1,
			StatusMsg: "token无效",
		}, nil
	}

	fileName := claims.Username + request.Title + fmt.Sprint(time.Now().Unix()) +".mp4"
	filePath := "static/" + fileName
	f, err := os.Create(filePath) 
	if err != nil {
		return &DouYinPublishActionResponse{
			StatusCode: -2,
			StatusMsg: "创建视频文件出错",
		}, err
	}
	defer f.Close()
	// 将视频字节数组写入文件
	if _, err := f.Write(request.Data); err != nil {
		return &DouYinPublishActionResponse{
			StatusCode: -2,
			StatusMsg: "视频字节数组写入文件出错",
		}, nil
	}
	// 将缓冲区的数据写入磁盘
	if err := f.Sync(); err != nil {
		return &DouYinPublishActionResponse{
			StatusCode: -2,
			StatusMsg: "缓冲区的数据写入磁盘出错",
		}, nil
	}
	// 保存数据至mysql
	video := model.Video{AuthorId: claims.ID, PlayUrl: filePath, Title: request.Title}
	_, err = model.CreateVideo(&video)
	if err != nil {
		return &DouYinPublishActionResponse{
			StatusCode: -2,
			StatusMsg: "保存视频出错",
		}, nil
	}

	//添加视频id至redis list
	go func ()  {
		err = model.AddVideoToList(video.Id)
		if err != nil {
			fmt.Println(err)
		}
	}()
	
	return &DouYinPublishActionResponse{
		StatusCode: 0,
		StatusMsg: "保存视频成功",
	}, nil

}

func (pL *publishListService) PublishList(context context.Context, request *DouYinPublishListRequest) (*DouYinPublishListResponse, error) {
	claims, _ := utils.ParseToken(request.Token)
	if (claims == nil) || (claims.ID != request.UserId) {
		return &DouYinPublishListResponse{
			StatusCode: -1,
			StatusMsg: "token无效",
			VideoList: nil,
		}, nil
	}

	videos, err := model.ListVideo(request.UserId)
	if err != nil {
		return &DouYinPublishListResponse{
			StatusCode: -2,
			StatusMsg: "获取视频列表失败",
			VideoList: nil,
		}, nil
	}

	var videoList []*Video
	for _, video := range *videos {
		user , err := model.FindUserById(&model.User{Id: video.AuthorId})
		if err != nil {
			return &DouYinPublishListResponse{
				StatusCode: -2,
				StatusMsg: "获取视频列表时用户信息查询失败",
				VideoList: nil,
			}, nil
		}
		videoList = append(videoList, &Video{
			Id: video.Id,
			Author: &User{
				Id: user.Id,
				Name: user.Username,
				FollowCount: user.FollowCount,
				FollowerCount: user.FollowerCount,
				IsFollow: user.IsFollow,
				Avatar: user.Avatar,
				BackgroundImage: user.BackgroundImage,
				Signature: user.Signature,
				TotalFavorited: user.TotalFavorited,
				WorkCount: user.WorkCount,
				FavoriteCount: user.FavoriteCount,
			},
			PlayUrl: video.PlayUrl,
			CoverUrl: video.CoverUrl,
			FavoriteCount: video.FavoriteCount,
			CommentCount: video.CommentCount,
			IsFavorite: video.IsFavorite,
			Title: video.Title,
		})
	}

	return &DouYinPublishListResponse{
		StatusCode: 0,
		StatusMsg: "获取视频列表成功",
		VideoList: videoList,
	}, nil

}

func (pA *publishActionService) mustEmbedUnimplementedPublishActionServiceServer() {}

func (pL *publishListService) mustEmbedUnimplementedPublishListServiceServer() {}