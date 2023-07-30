package model

import (
	_"errors"
	"github.com/hdbdn77/simplifiedTikTok/pkg/dao"
	_"gorm.io/gorm"
)

type Video struct {
    Id            int64  `gorm:"primaryKey;autoIncrement"` 
    AuthorId      int64  `gorm:"not null"`
    PlayUrl       string `gorm:"not null"`
    CoverUrl      string `gorm:"default:''"`
    FavoriteCount int64  `gorm:"default:0"`
    CommentCount  int64  `gorm:"default:0"`
    IsFavorite    bool   `gorm:"default:false"`
    Title         string `gorm:"default:'';size:20"`
}

func (Video) TableName() string {
    return "video" 
}

func CreateVideo(video *Video) (*Video, error) {
	// 获取数据库连接
	db := dao.GetDB()
	// 迁移模型
	db.AutoMigrate(&Video{})

	// 创建
	err := db.Create(video).Error
	return video, err
}

func ListVideo(authorId int64) (*[]Video, error) {
	// 获取数据库连接
	db := dao.GetDB()
	// 迁移模型
	db.AutoMigrate(&Video{})

	//查询
	var videos []Video
	err := db.Where("author_id = ?", authorId).Find(&videos).Error
	return &videos, err

}