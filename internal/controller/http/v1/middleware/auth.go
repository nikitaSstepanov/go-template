package middleware

import (
	"app/internal/entity/types"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gosuit/gins"
)

func (m *Middleware) CheckAccess(roles ...types.Role) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")

		if header == "" {
			gins.Abort(ctx, foundErr)
			return
		}

		parts := strings.Split(header, " ")
		if len(parts) < 2 {
			gins.Abort(ctx, bearerErr)
			return
		}

		bearer := parts[0]
		token := parts[1]

		if bearer != bearerType {
			gins.Abort(ctx, bearerErr)
			return
		}

		claims, err := m.auth.ValidateToken(token, false)
		if err != nil {
			gins.Abort(ctx, err)
			return
		}

		if len(roles) != 0 {
			if !slices.Contains(roles, claims.Role) {
				gins.Abort(ctx, forbiddenErr)
			}
		}

		ctx.Set("userId", claims.Id)

		ctx.Next()
	}
}
