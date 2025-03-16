package auth

import (
	"net/http"

	"app/internal/entity"

	"github.com/gin-gonic/gin"
	"github.com/gosuit/e"
	"github.com/gosuit/lec"
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
)

type AuthUseCase interface {
	Login(ctx lec.Context, user *entity.User) (*entity.Tokens, e.Error)
	Refresh(ctx lec.Context, refresh string) (*entity.Tokens, e.Error)
}

type Middleware interface {
	CheckAccess(roles ...string) gin.HandlerFunc
}
