package auth

import (
	"github.com/gin-gonic/gin"
	conv "github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/converter"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/dto"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/validator"
)

type Auth struct {
	usecase AuthUseCase
}

func New(uc AuthUseCase) *Auth {
	return &Auth{
		usecase: uc,
	}
}

func (a *Auth) Login(ctx *gin.Context) {
	var body dto.Login

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(badReq, badReqErr)
		return
	}

	if err := validator.Struct(body, validator.Password); err != nil {
		ctx.AbortWithStatusJSON(err.ToHttpCode(), err)
		return
	}

	user := conv.EntityLogin(body)

	tokens, err := a.usecase.Login(ctx, user)
	if err != nil {
		ctx.AbortWithStatusJSON(err.ToHttpCode(), err)
		return
	}

	ctx.SetCookie(cookieName, tokens.Refresh, cookieAge, cookiePath, cookieHost, false, true)

	result := conv.DtoToken(tokens)

	ctx.JSON(ok, result)
}

func (a *Auth) Logout(ctx *gin.Context) {
	userId := ctx.GetUint64("userId")

	err := a.usecase.Logout(ctx, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.ToHttpCode(), err)
		return
	}

	ctx.SetCookie(cookieName, "", -1, cookiePath, cookieHost, false, true)

	ctx.JSON(ok, logoutMsg)
}

func (a *Auth) Refresh(ctx *gin.Context) {
	refresh, cookieErr := ctx.Cookie(cookieName)
	if cookieErr != nil {
		ctx.AbortWithStatusJSON(unauth, unauthErr)
		return
	}

	tokens, err := a.usecase.Refresh(ctx, refresh)
	if err != nil {
		ctx.AbortWithStatusJSON(err.ToHttpCode(), err)
		return
	}

	ctx.SetCookie(cookieName, tokens.Refresh, cookieAge, cookiePath, cookieHost, false, true)

	result := conv.DtoToken(tokens)

	ctx.JSON(ok, result)
}
