package feedService
import (
	"context"
	"fmt"
	"time"
	"os"
	"errors"
	"encoding/json"

	"github.com/hdbdn77/simplifiedTikTok/pkg/model"
	"github.com/hdbdn77/simplifiedTikTok/pkg/utils"
	"github.com/hdbdn77/simplifiedTikTok/pkg/dao"
	"github.com/redis/go-redis/v9"
)

var FeedService = &feedService{}

type feedService struct {}

func (f *feedService) Feed(context context.Context, request *DouYinFeedRequest) (*DouYinFeedResponse, error) {
	//实现具体的业务逻辑
	// 检查token是否有效
	if request.Token != "" {
		claims, _ := utils.ParseToken(request.Token)
		if (claims == nil) {
			return &DouYinFeedResponse{
				StatusCode: -1,
				StatusMsg: "token无效",
				VideoList: nil,
				NextTime: time.Now().Unix(),
			}, nil
		}
	}

	if request.LatestTime == 0 {
		
	}
}