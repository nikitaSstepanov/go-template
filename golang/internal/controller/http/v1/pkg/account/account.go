package account

import (
	"github.com/gin-gonic/gin"
	conv "github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/converter"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/dto"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/validator"
)

type Account struct {
	usecase AccountUseCase
}

func New(uc AccountUseCase) *Account {
	return &Account{
		usecase: uc,
	}
}

func (a *Account) Get(c *gin.Context) {
	userId := c.GetUint64("userId")

	user, err := a.usecase.Get(c, userId)
	if err != nil {
		dto.AbortErrMsg(c, err)
		return
	}

	result := conv.DtoUser(user)

	c.JSON(ok, result)
}

func (a *Account) Create(c *gin.Context) {
	var body dto.CreateUser

	if err := c.ShouldBindJSON(&body); err != nil {
		dto.AbortErrMsg(c, badReqErr)
		return
	}

	if err := validator.Struct(body, validator.Password); err != nil {
		dto.AbortErrMsg(c, err)
		return
	}

	user := conv.EntityCreate(body)

	tokens, err := a.usecase.Create(c, user)
	if err != nil {
		dto.AbortErrMsg(c, err)
		return
	}

	c.SetCookie(cookieName, tokens.Refresh, cookieAge, cookiePath, cookieHost, false, true)

	result := conv.DtoToken(tokens)

	c.JSON(ok, result)
}

func (a *Account) Update(c *gin.Context) {
	userId := c.GetUint64("userId")

	var body dto.UpdateUser

	if err := c.ShouldBindJSON(&body); err != nil {
		dto.AbortErrMsg(c, badReqErr)
		return
	}

	if err := validator.Struct(body, validator.Password); err != nil {
		dto.AbortErrMsg(c, err)
		return
	}

	user := conv.EntityUpdate(body)
	user.Id = userId

	err := a.usecase.Update(c, user, body.OldPassword)
	if err != nil {
		dto.AbortErrMsg(c, err)
		return
	}

	c.JSON(ok, updatedMsg)
}

func (a *Account) Verify(c *gin.Context) {
	userId := c.GetUint64("userId")

	code := c.Param("code")

	if err := validator.StringLength(code, 6, 6); err != nil {
		dto.AbortErrMsg(c, err)
		return
	}

	err := a.usecase.Verify(c, userId, code)
	if err != nil {
		dto.AbortErrMsg(c, err)
		return
	}

	c.JSON(ok, verifiedMsg)
}

func (a *Account) ResendCode(c *gin.Context) {
	userId := c.GetUint64("userId")

	err := a.usecase.ResendCode(c, userId)
	if err != nil {
		dto.AbortErrMsg(c, err)
		return
	}

	c.JSON(ok, okMsg)
}

func (a *Account) Delete(c *gin.Context) {
	userId := c.GetUint64("userId")

	var body dto.DeleteUser

	if err := c.ShouldBindJSON(&body); err != nil {
		dto.AbortErrMsg(c, badReqErr)
		return
	}

	if err := validator.Struct(body, validator.Password); err != nil {
		dto.AbortErrMsg(c, err)
		return
	}

	user := conv.EntityDelete(body)
	user.Id = userId

	err := a.usecase.Delete(c, user)
	if err != nil {
		dto.AbortErrMsg(c, err)
		return
	}

	c.JSON(ok, okMsg)
}
