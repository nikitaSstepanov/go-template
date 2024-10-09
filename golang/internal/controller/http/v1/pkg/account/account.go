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

func (a *Account) Get(ctx *gin.Context) {
	userId := ctx.GetUint64("userId")

	user, err := a.usecase.Get(ctx, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.ToHttpCode(), err)
		return
	}

	result := conv.DtoUser(user)

	ctx.JSON(ok, result)
}

func (a *Account) Create(ctx *gin.Context) {
	var body dto.CreateUser

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(badReq, badReqErr)
		return
	}

	if err := validator.Struct(body, validator.Password); err != nil {
		ctx.AbortWithStatusJSON(err.ToHttpCode(), err)
		return
	}

	user := conv.EntityCreate(body)

	tokens, err := a.usecase.Create(ctx, user)
	if err != nil {
		ctx.AbortWithStatusJSON(err.ToHttpCode(), err)
		return
	}

	ctx.SetCookie(cookieName, tokens.Refresh, cookieAge, cookiePath, cookieHost, false, true)

	result := conv.DtoToken(tokens)

	ctx.JSON(ok, result)
}

func (a *Account) Update(ctx *gin.Context) {
	userId := ctx.GetUint64("userId")

	var body dto.UpdateUser

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(badReq, badReqErr)
		return
	}

	if err := validator.Struct(body, validator.Password); err != nil {
		ctx.AbortWithStatusJSON(err.ToHttpCode(), err)
		return
	}

	user := conv.EntityUpdate(body)
	user.Id = userId

	err := a.usecase.Update(ctx, user, body.OldPassword)
	if err != nil {
		ctx.AbortWithStatusJSON(err.ToHttpCode(), err)
		return
	}

	ctx.JSON(ok, updatedMsg)
}

func (a *Account) Verify(ctx *gin.Context) {
	userId := ctx.GetUint64("userId")

	code := ctx.Param("code")

	if err := validator.StringLength(code, 6, 6); err != nil {
		ctx.AbortWithStatusJSON(err.ToHttpCode(), err)
		return
	}

	err := a.usecase.Verify(ctx, userId, code)
	if err != nil {
		ctx.AbortWithStatusJSON(err.ToHttpCode(), err)
		return
	}

	ctx.JSON(ok, verifiedMsg)
}

func (a *Account) ResendCode(ctx *gin.Context) {
	userId := ctx.GetUint64("userId")

	err := a.usecase.ResendCode(ctx, userId)
	if err != nil {
		ctx.AbortWithStatusJSON(err.ToHttpCode(), err)
		return
	}

	ctx.JSON(ok, okMsg)
}

func (a *Account) Delete(ctx *gin.Context) {
	userId := ctx.GetUint64("userId")

	var body dto.DeleteUser

	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.AbortWithStatusJSON(badReq, badReqErr)
		return
	}

	if err := validator.Struct(body, validator.Password); err != nil {
		ctx.AbortWithStatusJSON(err.ToHttpCode(), err)
		return
	}

	user := conv.EntityDelete(body)
	user.Id = userId

	err := a.usecase.Delete(ctx, user)
	if err != nil {
		ctx.AbortWithStatusJSON(err.ToHttpCode(), err)
		return
	}

	ctx.JSON(ok, okMsg)
}
