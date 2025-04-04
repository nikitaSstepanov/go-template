package auth

import (
	"app/internal/entity"
	"app/internal/usecase/pkg/auth/jwt"
	"time"

	"github.com/gosuit/e"
	"github.com/gosuit/lec"
	"github.com/gosuit/utils/coder"
)

type UseCases struct {
	Jwt   JwtUseCase
	Coder coder.Coder
}

type Storages struct {
	User UserStorage
}

type JwtUseCase interface {
	ValidateToken(jwtString string, isRefresh bool) (*jwt.Claims, e.Error)
	GenerateToken(user *entity.User, expires time.Duration, isRefresh bool) (string, e.Error)
}

type UserStorage interface {
	GetById(ctx lec.Context, id uint64) (*entity.User, e.Error)
	GetByEmail(ctx lec.Context, email string) (*entity.User, e.Error)
}
