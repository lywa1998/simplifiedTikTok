package messageservice

import (
	"context"

	"github.com/micro/simplifiedTikTok/messageservice/pkg/utils"
)

var MessageChatService = &messageChatService{}
var MessageActionService = &messageActionService{}

type messageChatService struct{}
type messageActionService struct{}

func (mC *messageChatService) MessageChat(context context.Context, request *DouYinMessageChatRequest) (*DouYinMessageChatResponse, error) {
	if request.Token != "" {
		claims, _ := utils.ParseToken(request.Token)
		if claims == nil {
			return &DouYinMessageChatResponse{
				StatusCode:  -1,
				StatusMsg:   "token 无效",
				MessageList: nil,
			}, nil
		}
	}

	return &DouYinMessageChatResponse{}, nil
}

func (mA *messageActionService) MEssageAction(context context.Context, request *DouYinMessageActionRequest) (*DouYinMessageActionResponse, error) {

	return &DouYinMessageActionResponse{
        StatusCode: 0,
        StatusMsg: "消息发送成功",
    }, nil
}

func (mC *messageChatService) mustEmbedUnimplementedMessageChatServiceServer() {}

func (mA *messageActionService) mustEmbedUnimplementedMessageActionServiceServer() {}
