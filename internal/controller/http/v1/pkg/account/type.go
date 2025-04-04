package account

import (
	"app/internal/entity"

	"github.com/gosuit/e"
	"github.com/gosuit/lec"
)

type UserUseCase interface {
	Get(ctx lec.Context, userId uint64) (*entity.User, e.Error)
	Create(ctx lec.Context, user *entity.User) (*entity.Tokens, e.Error)
	Update(ctx lec.Context, user *entity.User, pass string) e.Error
	SetRole(ctx lec.Context, user *entity.User) e.Error
	Verify(ctx lec.Context, id uint64, code string) e.Error
	ResendCode(ctx lec.Context, userId uint64) e.Error
	Delete(ctx lec.Context, id uint64) e.Error
}
