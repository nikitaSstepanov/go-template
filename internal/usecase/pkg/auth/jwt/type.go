package jwt

import (
	"app/internal/entity/types"

	"github.com/golang-jwt/jwt/v5"
)

type JwtOptions struct {
	Audience   []string `yaml:"audience" env:"JWT_AUDIENCE"`
	Issuer     string   `yaml:"issuer"   env:"JWT_ISSUER"`
	AccessKey  string   `env:"JWT_ACCESS_SECRET"`
	RefreshKey string   `env:"JWT_REFRESH_SECRET"`
}

type Claims struct {
	Id   uint64     `json:"id"`
	Role types.Role `json:"role"`
	jwt.RegisteredClaims
}
