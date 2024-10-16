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

func (a *Auth) Login(c *gin.Context) {
	var body dto.Login

	if err := c.ShouldBindJSON(&body); err != nil {
		dto.AbortErrMsg(c, badReqErr)
		return
	}

	if err := validator.Struct(body, validator.Password); err != nil {
		dto.AbortErrMsg(c, err)
		return
	}

	user := conv.EntityLogin(body)

	tokens, err := a.usecase.Login(c, user)
	if err != nil {
		dto.AbortErrMsg(c, err)
		return
	}

	c.SetCookie(cookieName, tokens.Refresh, cookieAge, cookiePath, cookieHost, false, true)

	result := conv.DtoToken(tokens)

	c.JSON(ok, result)
}

func (a *Auth) Logout(c *gin.Context) {
	userId := c.GetUint64("userId")

	err := a.usecase.Logout(c, userId)
	if err != nil {
		dto.AbortErrMsg(c, err)
		return
	}

	c.SetCookie(cookieName, "", -1, cookiePath, cookieHost, false, true)

	c.JSON(ok, logoutMsg)
}

func (a *Auth) Refresh(c *gin.Context) {
	refresh, cookieErr := c.Cookie(cookieName)
	if cookieErr != nil {
		dto.AbortErrMsg(c, badReqErr)
		return
	}

	tokens, err := a.usecase.Refresh(c, refresh)
	if err != nil {
		dto.AbortErrMsg(c, err)
		return
	}

	c.SetCookie(cookieName, tokens.Refresh, cookieAge, cookiePath, cookieHost, false, true)

	result := conv.DtoToken(tokens)

	c.JSON(ok, result)
}
