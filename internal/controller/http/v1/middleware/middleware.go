package middleware

import "app/internal/usecase/pkg/auth"

type Middleware struct {
	jwt JwtUseCase
}

func New(jwtOpts *auth.JwtOptions) *Middleware {
	return &Middleware{
		jwt: auth.NewJwt(jwtOpts),
	}
}
