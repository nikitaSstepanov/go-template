package auth

import (
	"context"
	"time"

	"app/internal/entity"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gosuit/e"
)

const (
	refreshExpires = 72 * time.Hour
	accessExpires  = 1 * time.Hour
	cost           = 10
)

var (
	internalErr = e.New("Something going wrong...", e.Internal)
	badDataErr  = e.New("Incorrect email or password", e.Unauthorize)
	unauthErr   = e.New("Token is invalid", e.Unauthorize)
)

type JwtOptions struct {
	Audience   []string `yaml:"audience" env:"JWT_AUDIENCE"`
	Issuer     string   `yaml:"issuer"   env:"JWT_ISSUER"`
	AccessKey  string   `env:"JWT_ACCESS_SECRET"`
	RefreshKey string   `env:"JWT_REFRESH_SECRET"`
}

type Claims struct {
	Id   uint64 `json:"id"`
	Role string `json:"role"`
	jwt.RegisteredClaims
}

type UserStorage interface {
	GetById(ctx context.Context, id uint64) (*entity.User, e.Error)
	GetByEmail(ctx context.Context, email string) (*entity.User, e.Error)
}

type TokenStorage interface {
	Get(ctx context.Context, userId uint64) (*entity.Token, e.Error)
	Set(ctx context.Context, token *entity.Token) e.Error
	Del(ctx context.Context, userId uint64) e.Error
}
