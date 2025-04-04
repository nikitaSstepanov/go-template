package middleware

import "github.com/gosuit/e"

const (
	bearerType = "Bearer"
)

var (
	foundErr     = e.New("Authorization header wasn`t found.", e.Unauthorize)
	forbiddenErr = e.New("This resource is forbidden.", e.Forbidden)
	bearerErr    = e.New("Token is not bearer.", e.Unauthorize)
)
