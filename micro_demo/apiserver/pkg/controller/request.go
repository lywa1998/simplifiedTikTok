package controller

import (
	"mime/multipart"
)

type Request struct {
	UserId string `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

// user request
type UserLoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type UserInfoRequest struct {
	Request
}

// video request
type PublishRequest struct {
	Data  *multipart.FileHeader `form:"data" binding:"required"`
	Token string                `form:"token" binding:"required"`
	Title string                `form:"title" binding:"required"`
}

type PublishListRequest struct {
	Request
}

// feed request
type FeedRequest struct {
	LatestTime string `form:"latest_time"`
	Token      string `form:"token"`
}

// message request
type MessageChatRequest struct {
    Token string `form:"token" binding:"required"`
    ToUserId int64 `form:"to_user_id" binding:"required"`
    PreMsgTime int64 `form:"pre_msg_time" binding:"required"`
}

type MessageActionRequest struct {
    Token string  
    ToUserId int64
    ActionType int32
    Content string
}
