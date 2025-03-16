package middleware

import (
	"slices"
	"strings"

	"app/internal/usecase/pkg/auth"

	"github.com/gin-gonic/gin"
	"github.com/gosuit/e"
	"github.com/gosuit/gins"
)

const (
	bearerType = "Bearer"
)

var (
	foundErr     = e.New("Authorization header wasn`t found", e.Unauthorize)
	bearerErr    = e.New("Token is not bearer", e.Unauthorize)
	forbiddenErr = e.New("This resource is forbidden", e.Forbidden)
)

type JwtUseCase interface {
	ValidateToken(jwtString string, isRefresh bool) (*auth.Claims, e.Error)
}

func (m *Middleware) CheckAccess(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")

		if header == "" {
			gins.Abort(ctx,foundErr )
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

		claims, err := m.jwt.ValidateToken(token, false)
		if err != nil {
			gins.Abort(ctx, err)
			return
		}

		if len(roles) != 0 && !slices.Contains(roles, claims.Role) {
			gins.Abort(ctx, forbiddenErr)
			return
		}

		ctx.Set("userId", claims.Id)

		ctx.Next()
	}
}
