package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func (m *Middleware) InitLogger(log *slog.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := ctx.Request

		ctx.Next()
		entry := log.With(
			slog.String("method", req.Method),
			slog.String("path", req.URL.Path),
			slog.String("remote_addr", req.RemoteAddr),
			slog.String("user_agent", req.UserAgent()),
		)

		t1 := time.Now()
		defer func() {
			entry.Info("request completed",
				slog.Int("status", ctx.Writer.Status()),
				slog.String("duration", time.Since(t1).String()),
			)
		}()
	}
}
