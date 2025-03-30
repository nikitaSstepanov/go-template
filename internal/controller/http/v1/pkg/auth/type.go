package auth

import (
	"app/internal/entity"

	"github.com/gin-gonic/gin"
	"github.com/gosuit/e"
	"github.com/gosuit/lec"
)

type AuthUseCase interface {
	Login(ctx lec.Context, user *entity.User) (*entity.Tokens, e.Error)
	Refresh(ctx lec.Context, refresh string) (*entity.Tokens, e.Error)
}

type Middleware interface {
	CheckAccess(roles ...string) gin.HandlerFunc
}
