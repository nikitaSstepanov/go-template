package cl

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/middleware"
	"github.com/nikitaSstepanov/tools/sl"
)

func GetL(c *gin.Context) *slog.Logger {
	if logger, ok := c.Get(middleware.CtxLoggerKey); ok {
		return logger.(*slog.Logger)
	}

	return sl.Default()
}
