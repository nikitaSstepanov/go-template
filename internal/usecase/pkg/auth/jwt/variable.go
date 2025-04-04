package jwt

import "github.com/gosuit/e"

var (
	unauthErr = e.New("Token is invalid.", e.Unauthorize)
)
