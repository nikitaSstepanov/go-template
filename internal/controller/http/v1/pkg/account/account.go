package account

import (
	conv "app/internal/controller/http/v1/converter"
	"app/internal/controller/http/v1/dto"
	"app/internal/controller/http/v1/validator"

	"github.com/gin-gonic/gin"
	"github.com/gosuit/e"
	"github.com/gosuit/gins"
	"github.com/gosuit/httper"
)

type Account struct {
	usecase UserUseCase
	cookie  *httper.Cookie
}

func New(uc UserUseCase, cookie *httper.Cookie) *Account {
	return &Account{
		usecase: uc,
		cookie:  cookie,
	}
}

// @Summary Retrieve user own account
// @Description Returns user information.
// @Tags Account
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} dto.Account "Successful response"
// @Failure 401 {object} resp.JsonError "Authorization header wasn't found, Token is not bearer"
// @Failure 404 {object} resp.JsonError "This user wasn`t found."
// @Router /account [get]
func (a *Account) Get(c *gin.Context) {
	ctx := gins.GetCtx(c)
	userId := c.GetUint64("userId")

	user, err := a.usecase.Get(ctx, userId)
	if err != nil {
		gins.Abort(c, err)
		return
	}

	result := conv.DtoUser(user)

	c.JSON(httper.StatusOK, result)
}

// @Summary Create account
// @Description Creates a new account and returns access tokens.
// @Tags Account
// @Accept json
// @Produce json
// @Param body body dto.CreateUser true "Data for creating a user"
// @Success 200 {object} dto.Token "Successful response with token"
// @Failure 400 {object} resp.JsonError "Incorrect data"
// @Failure 403 {object} resp.JsonError "Incorrect password"
// @Failure 409 {object} resp.JsonError "User with this email already exist"
// @Router /account/new [post]
func (a *Account) Create(c *gin.Context) {
	ctx := gins.GetCtx(c)
	var body dto.CreateUser

	if err := c.ShouldBindJSON(&body); err != nil {
		gins.Abort(c, badReqErr.WithErr(err))
		return
	}

	if err := validator.Struct(body, validator.Password); err != nil {
		gins.Abort(c, err)
		return
	}

	user := conv.EntityCreate(body)

	tokens, err := a.usecase.Create(ctx, user)
	if err != nil {
		gins.Abort(c, err)
		return
	}

	c.SetCookie(
		a.cookie.Name, tokens.Refresh, a.cookie.Age, a.cookie.Path,
		a.cookie.Host, a.cookie.Secure, a.cookie.HttpOnly,
	)

	result := conv.DtoToken(tokens)

	c.JSON(httper.StatusOK, result)
}

// @Summary Update user information
// @Description Updates the user's information including password.
// @Tags Account
// @Accept json
// @Produce json
// @Param body body dto.UpdateUser true "User update data"
// @Security Bearer
// @Success 200 {object} resp.Message "Updated."
// @Failure 400 {object} resp.JsonError "Incorrect data."
// @Failure 401 {object} resp.JsonError "Authorization header wasn't found, Token is not bearer"
// @Failure 403 {object} resp.JsonError "This resource is forbidden"
// @Failure 404 {object} resp.JsonError "This user wasn't found"
// @Failure 409 {object} resp.JsonError "User with this email already exists"
// @Router /account/edit [patch]
func (a *Account) Update(c *gin.Context) {
	ctx := gins.GetCtx(c)
	userId := c.GetUint64("userId")

	var body dto.UpdateUser

	if err := c.ShouldBindJSON(&body); err != nil {
		gins.Abort(c, badReqErr.WithErr(err))
		return
	}

	if err := validator.Struct(body, validator.Password); err != nil {
		gins.Abort(c, err)
		return
	}

	user := conv.EntityUpdate(body)
	user.Id = userId

	err := a.usecase.Update(ctx, user, body.OldPassword)
	if err != nil {
		gins.Abort(c, err)
		return
	}

	c.JSON(httper.StatusOK, updatedMsg)
}

// @Summary Update user role
// @Description Updates the user's role. Available only for ADMIN
// @Tags Account
// @Accept json
// @Produce json
// @Param body body dto.SetRole true "User set role data"
// @Security Bearer
// @Success 200 {object} resp.Message "Updated."
// @Failure 400 {object} resp.JsonError "Incorrect data."
// @Failure 401 {object} resp.JsonError "Authorization header wasn't found, Token is not bearer"
// @Failure 403 {object} resp.JsonError "This resource is forbidden"
// @Failure 404 {object} resp.JsonError "This user wasn't found"
// @Router /account/edit [patch]
func (a *Account) SetRole(c *gin.Context) {
	ctx := gins.GetCtx(c)

	var body dto.SetRole

	if err := c.ShouldBindJSON(&body); err != nil {
		gins.Abort(c, e.BadInputErr.WithErr(err))
		return
	}

	if err := validator.Struct(body); err != nil {
		gins.Abort(c, err)
		return
	}

	user := conv.EntitySetRole(body)

	err := a.usecase.SetRole(ctx, user)
	if err != nil {
		gins.Abort(c, err)
		return
	}

	c.JSON(httper.StatusOK, updatedMsg)
}

// @Summary Verify user
// @Description Verifies user with the provided activation code.
// @Tags Account
// @Accept json
// @Produce json
// @Security Bearer
// @Param code path string true "Activation Code" minlength(6) maxlength(6)
// @Success 200 {object} resp.Message "Verified."
// @Failure 400 {object} resp.JsonError "Your activation code is wrong., Bad string length"
// @Failure 401 {object} resp.JsonError "Authorization header wasn`t found, Token is not bearer"
// @Failure 403 {object} resp.JsonError "This resource is forbidden"
// @Failure 404 {object} resp.JsonError "This code wasn`t found."
// @Router /account/edit/verify/confirm/{code} [patch]
func (a *Account) Verify(c *gin.Context) {
	ctx := gins.GetCtx(c)
	userId := c.GetUint64("userId")

	code := c.Param("code")

	if err := validator.StringLength(code, 6, 6); err != nil {
		gins.Abort(c, err)
		return
	}

	err := a.usecase.Verify(ctx, userId, code)
	if err != nil {
		gins.Abort(c, err)
		return
	}

	c.JSON(httper.StatusOK, verifiedMsg)
}

// @Summary Resend verification code
// @Description Resends a verification code to the user's email or phone number.
// @Tags Account
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} resp.Message "Ok."
// @Failure 400 {object} resp.JsonError "Incorrect data"
// @Failure 401 {object} resp.JsonError `"This resource is forbidden, Authorization header wasn`t found, Token is not bearer"`
// @Failure 403 {object} resp.JsonError "This resource is forbidden"
// @Failure 404 {object} resp.JsonError "User not found"
// @Router /account/edit/verify/resend [get]
func (a *Account) ResendCode(c *gin.Context) {
	ctx := gins.GetCtx(c)
	userId := c.GetUint64("userId")

	err := a.usecase.ResendCode(ctx, userId)
	if err != nil {
		gins.Abort(c, err)
		return
	}

	c.JSON(httper.StatusOK, okMsg)
}

// @Summary Delete user account
// @Description Deletes a user account.
// @Tags Account
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} resp.Message "Ok."
// @Failure 401 {object} resp.JsonError "Authorization header wasn`t found, Token is not bearer"
// @Router /account/delete [delete]
func (a *Account) Delete(c *gin.Context) {
	ctx := gins.GetCtx(c)

	userId := c.GetUint64("userId")

	err := a.usecase.Delete(ctx, userId)
	if err != nil {
		gins.Abort(c, err)
		return
	}

	c.JSON(httper.StatusNoContent, okMsg)
}
