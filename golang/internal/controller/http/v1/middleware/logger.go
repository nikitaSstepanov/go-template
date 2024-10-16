package middleware

import (
	"context"
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nikitaSstepanov/tools/sl"
)

func (m *Middleware) InitLogger(ctx context.Context) gin.HandlerFunc {
	log := sl.L(ctx)

	log.Info("logger middleware enabled.")
	return func(c *gin.Context) {
		req := c.Request

		c.Next()
		entry := log.With(
			slog.String("method", req.Method),
			slog.String("path", req.URL.Path),
			slog.String("remote_addr", req.RemoteAddr),
			slog.String("user_agent", req.UserAgent()),
		)

		t1 := time.Now()
		defer func() {
			entry.Info("request completed",
				slog.Int("status", c.Writer.Status()),
				slog.String("duration", time.Since(t1).String()),
			)
		}()
	}
}
