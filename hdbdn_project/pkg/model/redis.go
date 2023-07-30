package model

import(
	"errors"
	"context"

	"github.com/hdbdn77/simplifiedTikTok/pkg/dao"
	"github.com/redis/go-redis/v9"
)

func AddVideoToList(videoId int64) error {
	client := dao.GetClient()
	ctx := context.Background()
	// 检查list长度是否超过30
	len, err := client.LLen(ctx, "video").Result()
	if err == redis.Nil {
		// 列表不存在,先创建
		if _, err := client.RPush(ctx, "video", videoId).Result(); err != nil {
			return errors.New("update video list failed") 
		}
	}else if err != nil {
		return errors.New("update video list failed") 
	}else {
		// 如满了,就将左侧最早的删除掉
		if len >= 30 {
			if _, err := client.LPop(ctx, "video").Result(); err != nil {
				return errors.New("update video list failed")
			}
		}
		// 插入新数据
		if _, err := client.RPush(ctx, "video", videoId).Result(); err != nil {
			return errors.New("update video list failed")
		}

	}

	return nil
}