package model

import (
	"time"

	"github.com/micro/simplifiedTikTok/userservice/pkg/dao"
    "github.com/micro/simplifiedTikTok/messageservice/pkg/errno"
)

type Message struct {
    ID         int64 `gorm:"primaryKey;autoIncrement;comment:'PrimaryKey'"`
    ToUserId   int64 `gorm:"not null"`
    FromUserId int64 `gorm:"not null"`
    Content    string `gorm:"size"`
	CreateAt   time.Time
}

func (Message) TableName() string {
	return "messages"
}

func AddNewMessage(message *Message) error {
    exist, err := QueryUserById(message.FromUserId)
    if exist == nil || err != nil {
        return errno.UserIsNotExistErr
    }
    exist, err = QueryUserById(message.ToUserId)
    if exist == nil || err != nil {
        return errno.UserIsNotExistErr
    }

	db := dao.GetDB()
	err = db.Create(message).Error
    if err != nil {
        return err
    }
	return err
}

func GetMessageByIdPair(user_id1, user_id2 int64, pre_msg_time time.Time) ([]Message, error) {
    exist, err := QueryUserById(user_id1)
    if exist == nil || err != nil {
        return nil, errno.UserIsNotExistErr
    }
    exist, err = QueryUserById(user_id2)
    if exist == nil || err != nil {
        return nil, errno.UserIsNotExistErr
    }

    db := dao.GetDB()
    var messages []Message
    err = db.Where("to_user_id = ? AND from_user_id = ? AND created_at > ?", user_id1, user_id2, pre_msg_time).Error
    if err != nil {
        return nil, err
    }
    return messages, nil
}

func GetLatestMessageByIdPair(user_id1, user_id2 int64) (*Message, error) {
    exist, err := QueryUserById(user_id1)
    if exist == nil || err != nil {
        return nil, errno.UserIsNotExistErr
    }
    exist, err = QueryUserById(user_id2)
    if exist == nil || err != nil {
        return nil, errno.UserIsNotExistErr
    }

    db := dao.GetDB()
    var message Message
    err = db.Where("to_user_id = ? AND from_user_id = ?", user_id1, user_id2).Or("to_user_id = ? AMD from_user_id = ?", user_id2, user_id1).Last(&message).Error

    if err != nil {
        return nil, err
    }
    return &message, nil
}
