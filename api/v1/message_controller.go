package v1

import (
	"net/http"

	"chat-room/pkg/global/log"
	"chat-room/pkg/common/request"
	"chat-room/pkg/common/response"
	"chat-room/internal/service"

	"github.com/gin-gonic/gin"
)

// 获取消息列表
func GetMessage(c *gin.Context) {
	log.Info(c.Query("uuid"))
	var messageRequest request.MessageRequest
	err := c.BindQuery(&messageRequest)
	if nil != err {
		log.Error("bindQueryError", log.Any("bindQueryError", err))
	}
	log.Info("messageRequest params: ", log.Any("messageRequest", messageRequest))

	messages, err := service.MessageService.GetMessages(messageRequest)
	if err != nil {
		c.JSON(http.StatusOK, response.FailMsg(err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.SuccessMsg(messages))
}