package middleware

import (
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/dto"
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
			dto.AbortErrMsg(ctx, foundErr)
			return
		}

		parts := strings.Split(header, " ")
		bearer := parts[0]
		token := parts[1]

		if bearer != bearerType {
			dto.AbortErrMsg(ctx, bearerErr)
			return
		}

		claims, err := m.jwt.ValidateToken(token, false)
		if err != nil {
			dto.AbortErrMsg(ctx, err)
			return
		}

		if len(roles) != 0 && !slices.Contains(roles, claims.Role) {
			dto.AbortErrMsg(ctx, forbiddenErr)
			return
		}

		ctx.Set("userId", claims.Id)

		ctx.Next()
	}
}
