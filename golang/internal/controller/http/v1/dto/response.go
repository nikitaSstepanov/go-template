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

func AbortErrMsg(c *gin.Context, err e.Error) {
	cl.GetL(c).Error("Something going wrong", err.SlErr())

	c.AbortWithStatusJSON(
		err.ToHttpCode(),
		err.ToJson(),
	)
}
