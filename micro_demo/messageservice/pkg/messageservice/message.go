package messageservice

import (
	"context"

	"github.com/micro/simplifiedTikTok/messageservice/pkg/model"
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

    messages := make([]*Message, 0)
    token := request.GetToken()
    to_user_id := request.GetToUserId()
    pre_msg_time := request.GetPreMsgTime()

    db_messages, err := model.GetMessageByIdPair(utils.ParseToID(token), to_user_id, utils.MillTimeStampToTime(pre_msg_time))
    if err != nil {
        return &DouYinMessageChatResponse{
            StatusCode: 0,
            StatusMsg: "获取消息成功",
            MessageList: nil,
        }, nil
    }

    for _, db_message := range db_messages {
        messages = append(messages, &Message{
            Id: db_message.ID,
            ToUserId: db_message.ToUserId,
            FromUserId: db_message.FromUserId,
            Content: db_message.Content,
            CreateTime: db_message.CreateAt.UnixNano() / 1000000,
        })
    }

	return &DouYinMessageChatResponse{
        StatusCode: 0,
        StatusMsg: "获取消息成功",
        MessageList: messages,
    }, nil
}

func (mA *messageActionService) MessageAction(context context.Context, request *DouYinMessageActionRequest) (*DouYinMessageActionResponse, error) {
    token := request.GetToken()
    to_user_id := request.GetToUserId()
    content := request.GetContent()

    err := model.AddNewMessage(&model.Message{
        FromUserId: utils.ParseToID(token), 
        ToUserId: to_user_id, 
        Content: content, 
    })
    if err != nil {
        return &DouYinMessageActionResponse{
            StatusCode: -1,
            StatusMsg: "消息发送失败",
        }, err
    }

	return &DouYinMessageActionResponse{
        StatusCode: 0,
        StatusMsg: "消息发送成功",
    }, nil
}

func (mC *messageChatService) mustEmbedUnimplementedMessageChatServiceServer() {}

func (mA *messageActionService) mustEmbedUnimplementedMessageActionServiceServer() {}
