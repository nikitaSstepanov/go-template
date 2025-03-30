package middleware

import (
	"app/internal/usecase/pkg/auth"

	"github.com/gosuit/e"
)

type AuthUseCase interface {
	ValidateToken(jwtString string, isRefresh bool) (*auth.Claims, e.Error)
}
