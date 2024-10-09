package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/nikitaSstepanov/templates/golang/internal/usecase/pkg/auth"
	e "github.com/nikitaSstepanov/tools/error"
)

const (
	bearerType = "Bearer"
)

var (
	foundErr     = e.New("Authorization header wasn`t found", e.Unauthorize)
	bearerErr    = e.New("Token is not bearer", e.Unauthorize)
	forbiddenErr = e.New("This resource is forbidden", e.Unauthorize)
)

type JwtUseCase interface {
	ValidateToken(jwtString string, isRefresh bool) (*auth.Claims, e.Error)
}

func (m *Middleware) CheckAccess(roles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		header := ctx.GetHeader("Authorization")

		if header == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, foundErr)
			return
		}

		parts := strings.Split(header, " ")
		bearer := parts[0]
		token := parts[1]

		if bearer != bearerType {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, bearerErr)
			return
		}

		claims, err := m.jwt.ValidateToken(token, false)
		if err != nil {
			ctx.AbortWithStatusJSON(err.ToHttpCode(), err)
			return
		}

		if len(roles) != 0 && !slices.Contains(roles, claims.Role) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, forbiddenErr)
			return
		}

		ctx.Set("userId", claims.Id)

		ctx.Next()
	}
}
