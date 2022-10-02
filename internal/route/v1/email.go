package v1

import (
	"com.caiflower/commons/thirdpart/internal/config"
	"com.caiflower/commons/thirdpart/internal/entity"
	"com.caiflower/commons/thirdpart/internal/utils/mail"
	"github.com/gin-gonic/gin"
)

func send(message *entity.EmailByteMessage) error {
	auth := mail.NewAuthLogin(config.Config.UserName, config.Config.Password)
	client, e := auth.GetClient()
	if e != nil {
		return e
	}

	if err := client.SendMsg(message); err != nil {
		return err
	}

	client.Close()
	return nil
}

func SendString(c *gin.Context) {
	p := new(entity.EmailStringMessage)
	if e := c.ShouldBindJSON(p); e != nil {
		sendFailResponse(c, e)
		return
	}

	msg := new(entity.EmailByteMessage)
	msg.From = p.From
	msg.To = p.To
	msg.Title = p.Title
	msg.Content = []byte(p.Content)
	msg.ContentType = p.Content
	msg.Attachment = p.Attachment

	if e := send(msg); e != nil {
		sendFailResponse(c, e)
		return
	}

	sendSuccessResponse(c, entity.Response{Code: 0, Msg: "发送成功"})
}

func SendByte(c *gin.Context) {
	p := new(entity.EmailByteMessage)
	if e := c.ShouldBindJSON(p); e != nil {
		sendFailResponse(c, e)
		return
	}

	if e := send(p); e != nil {
		sendFailResponse(c, e)
		return
	}

	sendSuccessResponse(c, entity.Response{Code: 0, Msg: "发送成功"})
}
