package controller

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/micro/simplifiedTikTok/apiserver/pkg/clientconnect"
	"github.com/micro/simplifiedTikTok/apiserver/pkg/feedservice"
)

func Feed(c *gin.Context) {
	var feedRequest FeedRequest
	err := c.ShouldBindQuery(&feedRequest)
	if err != nil {
		c.JSON(http.StatusOK, FeedResponse{
			Response : Response{
				StatusCode: -1, 
				StatusMsg: "视频流信息输入有误",
			},
		})
	}

	latestTime, _ := strconv.ParseInt(feedRequest.LatestTime, 10, 64)
	feedServiceClient := <- clientconnect.FeedChan
	feedResponse, err := feedServiceClient.Feed(context.Background(), &feedservice.DouYinFeedRequest{LatestTime: latestTime, Token: feedRequest.Token})
	clientconnect.FeedChan <- feedServiceClient

	if (feedResponse == nil) || (err != nil) {
		fmt.Println(err)
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: -1,
				StatusMsg: "feed failed",
			},
		})
		return
	}
	if feedResponse.StatusCode != 0 {
		c.JSON(http.StatusOK, FeedResponse{
			Response: Response{
				StatusCode: feedResponse.StatusCode,
				StatusMsg: feedResponse.StatusMsg,
			},
			NextTime: feedResponse.NextTime,
		})
		return
	}
	var videoList []Video
	for _, video := range feedResponse.VideoList {
		videoList = append(videoList, Video{
			Id: video.Id,
			Author: User{
				Id: video.Author.Id,
				Name: video.Author.Name,
				FollowCount: video.Author.FollowCount,
				FollowerCount: video.Author.FollowerCount,
				IsFollow: video.Author.IsFollow,
				Avatar: video.Author.Avatar,
				BackgroundImage: video.Author.BackgroundImage,
				Signature: video.Author.Signature,
				TotalFavorited: video.Author.TotalFavorited,
				WorkCount: video.Author.WorkCount,
				FavoriteCount: video.Author.FavoriteCount,
			},
			PlayUrl: "https://www.w3schools.com/html/movie.mp4",
			CoverUrl: "http://5b0988e595225.cdn.sohucs.com/images/20180430/fcf555aed1804ad586b24b3aeda6c031.jpeg",
			FavoriteCount: video.FavoriteCount,
			CommentCount: video.CommentCount,
			IsFavorite: video.IsFavorite,
			Title: video.Title,
		})
	}
	c.JSON(http.StatusOK, FeedResponse{
		Response: Response{
			StatusCode: feedResponse.StatusCode,
			StatusMsg: feedResponse.StatusMsg,
		},
		NextTime: feedResponse.NextTime,
		VideoList: videoList,
	})
}