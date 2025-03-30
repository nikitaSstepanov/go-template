package auth

import (
	"app/internal/entity"

	"github.com/golang-jwt/jwt/v5"
	"github.com/gosuit/e"
	"github.com/gosuit/lec"
	"github.com/gosuit/utils/coder"
)

type UseCases struct {
	Jwt   *Jwt
	Coder coder.Coder
}

type Storages struct {
	User UserStorage
}

type JwtOptions struct {
	Audience   []string `yaml:"audience" env:"JWT_AUDIENCE"`
	Issuer     string   `yaml:"issuer"   env:"JWT_ISSUER"`
	AccessKey  string   `env:"JWT_ACCESS_SECRET"`
	RefreshKey string   `env:"JWT_REFRESH_SECRET"`
}

type Claims struct {
	Id uint64 `json:"id"`
	jwt.RegisteredClaims
}

type UserStorage interface {
	GetById(ctx lec.Context, id uint64) (*entity.User, e.Error)
	GetByEmail(ctx lec.Context, email string) (*entity.User, e.Error)
}
