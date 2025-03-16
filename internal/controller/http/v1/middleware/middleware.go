package middleware

import "app/internal/usecase/pkg/auth"

type Middleware struct {
	auth AuthUseCase
}

func New(uc *auth.Auth) *Middleware {
	return &Middleware{
		auth: uc,
	}
}
