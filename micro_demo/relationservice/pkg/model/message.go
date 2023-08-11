package model

import (
	"time"

	"github.com/micro/simplifiedTikTok/userservice/pkg/dao"
)

type Message struct {
	ID         int64
	ToUserId   int64
	FromUserId int64
	Content    string
	CreateAt   time.Time
}

func (Message) TableName() string {
	return "messages"
}

func AddNewMessage(message *Message) error {
	db := dao.GetDB()

	err := db.Create(message).Error
	return err
}
