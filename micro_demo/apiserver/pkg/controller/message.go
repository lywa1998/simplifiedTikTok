package controller

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/simplifiedTikTok/apiserver/pkg/clientconnect"
	"github.com/micro/simplifiedTikTok/messageservice/pkg/messageservice"
)

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	var messageActionRequest MessageActionRequest
	err := c.ShouldBindQuery(&messageActionRequest)
	if err != nil {
		c.JSON(http.StatusOK, MessageActionResponse{
            Response: Response{
                StatusCode: -1,
                StatusMsg: "消息发送失败",
            },
        })
	}

    messageServiceClient := <- clientconnect.MessageActionChan
    messageActionResponse, err := messageServiceClient.MessageAction(context.Background(), &messageservice.DouYinMessageActionRequest{
        Token: messageActionRequest.Token,
        ToUserId: messageActionRequest.ToUserId,
        ActionType: messageActionRequest.ActionType,
        Content: messageActionRequest.Content,

    })
    clientconnect.MessageActionChan <- messageServiceClient

    if (messageActionResponse == nil) || (err != nil) {
		fmt.Println(err)
		c.JSON(http.StatusOK, MessageActionResponse{
			Response: Response{
				StatusCode: -1,
				StatusMsg: "register err",
			},
		})
		return
	}
	if messageActionResponse.StatusCode != 0 {
		c.JSON(http.StatusOK, MessageActionResponse{
			Response: Response{
				StatusCode: messageActionResponse.StatusCode,
				StatusMsg: messageActionResponse.StatusMsg,
			},
		})
		return
	}
    c.JSON(http.StatusOK, MessageActionResponse{
        Response: Response{
            StatusCode: 0,
            StatusMsg: "消息发送成功",
        },
    })
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
    var messageChatRequest MessageChatRequest
	err := c.ShouldBindQuery(&messageChatRequest)
	if err != nil {
		c.JSON(http.StatusOK, MessageChatResponse{
            Response: Response{
                StatusCode: -1,
                StatusMsg: "消息发送失败",
            },
        })
	}

    messageServiceClient := <- clientconnect.MessageChatChan
    messageChatResponse, err := messageServiceClient.MessageChat(context.Background(), &messageservice.DouYinMessageChatRequest{
        Token: messageChatRequest.Token,
        ToUserId: messageChatRequest.ToUserId,
        PreMsgTime: messageChatRequest.PreMsgTime,
    })
    clientconnect.MessageChatChan <- messageServiceClient

    c.JSON(http.StatusOK, messageChatResponse)
}

func genChatKey(userIdA int64, userIdB int64) string {
	if userIdA > userIdB {
		return fmt.Sprintf("%d_%d", userIdB, userIdA)
	}
	return fmt.Sprintf("%d_%d", userIdA, userIdB)
}
