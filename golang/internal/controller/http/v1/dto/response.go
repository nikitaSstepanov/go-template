package dto

import (
	"github.com/gin-gonic/gin"
	cl "github.com/nikitaSstepanov/templates/golang/pkg/utils/controller"
	e "github.com/nikitaSstepanov/tools/error"
)

type Message struct {
	Message string `json:"message"`
}

func NewMessage(msg string) *Message {
	return &Message{
		Message: msg,
	}
}

type ErrMessage struct {
	Msg string `json:"error"`
}

func AbortErrMsg(c *gin.Context, err e.Error) {
	cl.GetL(c).Error(err.Error())

	c.AbortWithStatusJSON(
		err.ToHttpCode(),
		&ErrMessage{Msg: err.GetMessage()},
	)
}
