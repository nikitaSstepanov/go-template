package middleware

import (
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	resp "github.com/nikitaSstepanov/templates/golang/internal/controller/response"
	"github.com/nikitaSstepanov/templates/golang/internal/usecase/pkg/auth"
	e "github.com/nikitaSstepanov/tools/error"
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
			resp.AbortErrMsg(ctx, foundErr)
			return
		}

		parts := strings.Split(header, " ")
		if len(parts) < 2 {
			resp.AbortErrMsg(ctx, bearerErr)
			return
		}
		bearer := parts[0]
		token := parts[1]

		if bearer != bearerType {
			resp.AbortErrMsg(ctx, bearerErr)
			return
		}

		claims, err := m.jwt.ValidateToken(token, false)
		if err != nil {
			resp.AbortErrMsg(ctx, err)
			return
		}

		if len(roles) != 0 && !slices.Contains(roles, claims.Role) {
			resp.AbortErrMsg(ctx, forbiddenErr)
			return
		}

		ctx.Set("userId", claims.Id)

		ctx.Next()
	}
}
