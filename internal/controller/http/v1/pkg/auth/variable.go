package auth

import (
	"context"
	"net/http"

	resp "app/internal/controller/response"
	"app/internal/entity"

	"github.com/gin-gonic/gin"
	"github.com/gosuit/e"
)

const (
	ok      = http.StatusOK
	created = http.StatusCreated
	badReq  = http.StatusBadRequest
	unauth  = http.StatusUnauthorized

	cookieName = "refreshToken"
	cookieAge  = 259200
	cookiePath = "/"
	cookieHost = "localhost"
)

var (
	badReqErr = e.New("Incorrect data.", e.BadInput)

	logoutMsg = resp.NewMessage("Logout success.")
)

type AuthUseCase interface {
	Login(ctx context.Context, user *entity.User) (*entity.Tokens, e.Error)
	Logout(ctx context.Context, userId uint64) e.Error
	Refresh(ctx context.Context, refresh string) (*entity.Tokens, e.Error)
}

type Middleware interface {
	CheckAccess(roles ...string) gin.HandlerFunc
}
