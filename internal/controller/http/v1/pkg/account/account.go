package account

import (
	"github.com/gin-gonic/gin"
	conv "app/internal/controller/http/v1/converter"
	"app/internal/controller/http/v1/dto"
	"app/internal/controller/http/v1/validator"
	resp "app/internal/controller/response"
)

type Account struct {
	usecase AccountUseCase
}

func New(uc AccountUseCase) *Account {
	return &Account{
		usecase: uc,
	}
}

// @Summary Retrieve user by ID
// @Description Returns user information based on their ID.
// @Tags account
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} dto.Account "Successful response"
// @Failure 404 {object} dto.JsonError "This user wasn`t found."
// @Failure 500 {object} dto.JsonError "Something going wrong..."
// @Router /account/ [get]
func (a *Account) Get(c *gin.Context) {
	userId := c.GetUint64("userId")

	user, err := a.usecase.Get(c, userId)
	if err != nil {
		resp.AbortErrMsg(c, err)
		return
	}

	result := conv.DtoUser(user)

	c.JSON(ok, result)
}

// @title Create User
// @Summary Create User
// @Description Creates a new user and returns access tokens.
// @Tags account
// @Accept json
// @Produce json
// @Param body body dto.CreateUser true "Data for creating a user"
// @Success 200 {object} dto.Token "Successful response with token"
// @Failure 400 {object} dto.JsonError "Incorrect data"
// @Failure 403 {object} dto.JsonError "Incorrect password"
// @Failure 409 {object} dto.JsonError "User with this email already exist"
// @Failure 500 {object} dto.JsonError "Something going wrong..."
// @Router /account/new [post]
func (a *Account) Create(c *gin.Context) {
	var body dto.CreateUser

	if err := c.ShouldBindJSON(&body); err != nil {
		resp.AbortErrMsg(c, badReqErr.WithErr(err))
		return
	}

	if err := validator.Struct(body, validator.Password); err != nil {
		resp.AbortErrMsg(c, err)
		return
	}

	user := conv.EntityCreate(body)

	tokens, err := a.usecase.Create(c, user)
	if err != nil {
		resp.AbortErrMsg(c, err)
		return
	}

	c.SetCookie(cookieName, tokens.Refresh, cookieAge, cookiePath, cookieHost, false, true)

	result := conv.DtoToken(tokens)

	c.JSON(ok, result)
}

// @Summary Update user information
// @Description Updates the user's information including password.
// @Tags account
// @Accept json
// @Produce json
// @Param body body dto.UpdateUser true "User update data"
// @Security Bearer
// @Success 200 {object} dto.Message "Updated."
// @Failure 400 {object} dto.JsonError "Incorrect data., Your activation code is wrong."
// @Failure 400 {object} dto.JsonError
// @Failure 401 {object} dto.JsonError "Authorization header wasn't found, Token is not bearer"
// @Failure 403 {object} dto.JsonError "This resource is forbidden"
// @Failure 404 {object} dto.JsonError "This user wasn't found"
// @Failure 409 {object} dto.JsonError "User with this email already exists"
// @Failure 500 {object} dto.JsonError "Something going wrong..."
// @Router /account/edit [patch]
func (a *Account) Update(c *gin.Context) {
	userId := c.GetUint64("userId")

	var body dto.UpdateUser

	if err := c.ShouldBindJSON(&body); err != nil {
		resp.AbortErrMsg(c, badReqErr.WithErr(err))
		return
	}

	if err := validator.Struct(body, validator.Password); err != nil {
		resp.AbortErrMsg(c, err)
		return
	}

	user := conv.EntityUpdate(body)
	user.Id = userId

	err := a.usecase.Update(c, user, body.OldPassword)
	if err != nil {
		resp.AbortErrMsg(c, err)
		return
	}

	c.JSON(ok, updatedMsg)
}

// @Summary Verify user activation code
// @Description Verifies the provided activation code for the user.
// @Tags account
// @Accept json
// @Produce json
// @Security Bearer
// @Param code path string true "Activation Code" minlength(6) maxlength(6)
// @Success 200 {object} dto.Message "Verified."
// @Failure 400 {object} dto.JsonError "Your activation code is wrong., Bad string length"
// @Failure 401 {object} dto.JsonError "Authorization header wasn`t found, Token is not bearer"
// @Failure 403 {object} dto.JsonError "This resource is forbidden"
// @Failure 404 {object} dto.JsonError "This code wasn`t found."
// @Failure 500 {object} dto.JsonError "Something going wrong..."
// @Router /account/verify/confirm/{code} [get]
func (a *Account) Verify(c *gin.Context) {
	userId := c.GetUint64("userId")

	code := c.Param("code")

	if err := validator.StringLength(code, 6, 6); err != nil {
		resp.AbortErrMsg(c, err)
		return
	}

	err := a.usecase.Verify(c, userId, code)
	if err != nil {
		resp.AbortErrMsg(c, err)
		return
	}

	c.JSON(ok, verifiedMsg)
}

// @Summary Resend verification code
// @Description Resends a verification code to the user's email or phone number.
// @Tags account
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} dto.Message "Ok."
// @Failure 400 {object} dto.JsonError "Incorrect data"
// @Failure 401 {object} dto.JsonError `"This resource is forbidden, Authorization header wasn`t found, Token is not bearer"`
// @Failure 403 {object} dto.JsonError "This resource is forbidden"
// @Failure 404 {object} dto.JsonError "User not found"
// @Failure 500 {object} dto.JsonError "Something going wrong..."
// @Router /account/verify/resend [get]
func (a *Account) ResendCode(c *gin.Context) {
	userId := c.GetUint64("userId")

	err := a.usecase.ResendCode(c, userId)
	if err != nil {
		resp.AbortErrMsg(c, err)
		return
	}

	c.JSON(ok, okMsg)
}

// @Summary Delete user account
// @Description Deletes a user account by ID.
// @Tags account
// @Accept json
// @Produce json
// @Param body body dto.DeleteUser true "Delete User Request"
// @Security Bearer
// @Success 200 {object} dto.Message "Ok."
// @Failure 400 {object} dto.JsonError "Incorrect data"
// @Failure 401 {object} dto.JsonError "Authorization header wasn`t found, Token is not bearer"
// @Failure 403 {object} dto.JsonError "This resource is forbidden"
// @Failure 500 {object} dto.JsonError "Something going wrong..."
// @Router /account/delete [delete]
func (a *Account) Delete(c *gin.Context) {
	userId := c.GetUint64("userId")

	var body dto.DeleteUser

	if err := c.ShouldBindJSON(&body); err != nil {
		resp.AbortErrMsg(c, badReqErr.WithErr(err))
		return
	}

	if err := validator.Struct(body, validator.Password); err != nil {
		resp.AbortErrMsg(c, err)
		return
	}

	user := conv.EntityDelete(body)
	user.Id = userId

	err := a.usecase.Delete(c, user)
	if err != nil {
		resp.AbortErrMsg(c, err)
		return
	}

	c.JSON(ok, okMsg)
}
