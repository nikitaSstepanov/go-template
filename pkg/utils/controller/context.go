package cl

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/nikitaSstepanov/tools/sl"
)

const (
	CtxLoggerKey = "log"
)

func GetL(c *gin.Context) *slog.Logger {
	if logger, ok := c.Get(CtxLoggerKey); ok {
		return logger.(*slog.Logger)
	}

	return sl.Default()
}
