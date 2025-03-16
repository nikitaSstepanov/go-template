package resp

import (
	"github.com/gin-gonic/gin"
	cl "app/pkg/utils/controller"
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

// JsonError use only for doc and represent e.JsonError
type JsonError struct {
	Error string `json:"error"`
}

func AbortErrMsg(c *gin.Context, err e.Error) {
	log := cl.GetL(c)

	if err.GetCode() == e.Internal {
		log.Error("Something going wrong...", err.SlErr())
	} else {
		log.Info("Invalid input data", err.SlErr())
	}

	c.AbortWithStatusJSON(
		err.ToHttpCode(),
		err.ToJson(),
	)
}
