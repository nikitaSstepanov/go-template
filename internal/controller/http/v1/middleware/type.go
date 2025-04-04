package middleware

import (
	"app/internal/usecase/pkg/auth/jwt"

	"github.com/gosuit/e"
)

type AuthUseCase interface {
	ValidateToken(jwtString string, isRefresh bool) (*jwt.Claims, e.Error)
}
