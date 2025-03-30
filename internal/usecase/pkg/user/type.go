package user

import (
	"app/internal/entity"
	"time"

	"github.com/gosuit/e"
	"github.com/gosuit/lec"
	"github.com/gosuit/utils/coder"
)

type UseCases struct {
	Jwt   JwtUseCase
	Mail  MailUseCase
	Coder coder.Coder
}

type Storages struct {
	User UserStorage
	Code CodeStorage
}

type UserStorage interface {
	GetById(ctx lec.Context, id uint64) (*entity.User, e.Error)
	GetByEmail(ctx lec.Context, email string) (*entity.User, e.Error)
	Create(ctx lec.Context, user *entity.User) e.Error
	Update(ctx lec.Context, user *entity.User) e.Error
	Verify(ctx lec.Context, user *entity.User) e.Error
	Delete(ctx lec.Context, user *entity.User) e.Error
}

type CodeStorage interface {
	Get(ctx lec.Context, userId uint64) (*entity.ActivationCode, e.Error)
	Set(ctx lec.Context, code *entity.ActivationCode) e.Error
	Del(ctx lec.Context, userId uint64) e.Error
}

type JwtUseCase interface {
	GenerateToken(user *entity.User, expires time.Duration, isRefresh bool) (string, e.Error)
}

type MailUseCase interface {
	SendActivation(to string, code string) e.Error
}
