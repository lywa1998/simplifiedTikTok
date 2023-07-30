package model

import (
	"errors"
	"github.com/hdbdn77/simplifiedTikTok/pkg/dao"
	"gorm.io/gorm"
)

type User struct {
	Id              int64  `gorm:"primaryKey;autoIncrement;comment:'PrimaryKey'"`
	Username        string `gorm:"size:32;not null;default:'';comment:'Username'"`
	Password        string `gorm:"size:32;not null;default:'';comment:'Password'"`
	FollowCount     int64  `gorm:"not null;default:0;comment:'FollowCount'"`
	FollowerCount   int64  `gorm:"not null;default:0;comment:'FollowerCount'"`
	IsFollow        bool   `gorm:"not null;default:false;comment:'IsFollow'"`
	Avatar          string `gorm:"size:128;not null;default:'';comment:'Avatar'"`
	BackgroundImage string `gorm:"size:128;not null;default:'';comment:'BackgroundImage'"`
	Signature       string `gorm:"size:256;not null;default:'';comment:'Signature'"`
	TotalFavorited  int64  `gorm:"not null;default:0;comment:'TotalFavorited'"`
	WorkCount       int64  `gorm:"not null;default:0;comment:'WorkCount'"`
	FavoriteCount   int64  `gorm:"not null;default:0;comment:'FavoriteCount'"`
}

// 用户名索引
func (User) TableName() string {
	return "user"
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	// 自定义Username唯一性校验
	var count int64
	tx.Model(u).Where("username = ?", u.Username).Count(&count)
	if count > 0 {
		return errors.New("username already existed")
	}

	return nil
}

func Register(user *User) (*User, error) {
	// 获取数据库连接
	db := dao.GetDB()
	// 迁移模型
	db.AutoMigrate(&User{})

	// 创建
	err := db.Create(user).Error
	return user, err
}

func FindUserByUsername(user *User) (*User, error) {
	// 获取数据库连接
	db := dao.GetDB()
	// 迁移模型
	db.AutoMigrate(&User{})

	// 查询
	err := db.Where("username = ?", user.Username).Take(&user).Error
	return user, err
}
