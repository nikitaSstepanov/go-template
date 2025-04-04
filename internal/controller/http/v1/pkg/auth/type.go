package auth

import (
	"app/internal/entity"

	"github.com/gosuit/e"
	"github.com/gosuit/lec"
)

type AuthUseCase interface {
	Login(ctx lec.Context, user *entity.User) (*entity.Tokens, e.Error)
	Refresh(ctx lec.Context, refresh string) (*entity.Tokens, e.Error)
}
