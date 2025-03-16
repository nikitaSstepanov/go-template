package auth

import (
	conv "app/internal/controller/http/v1/converter"
	"app/internal/controller/http/v1/dto"
	"app/internal/controller/http/v1/validator"

	"github.com/gin-gonic/gin"
	"github.com/gosuit/gins"
)

type Auth struct {
	usecase AuthUseCase
}

func New(uc AuthUseCase) *Auth {
	return &Auth{
		usecase: uc,
	}
}

// @Summary Log in a user
// @Description Logs in a user with email and password
// @Tags auth
// @Accept json
// @Produce json
// @Param body body dto.Login true "Login information"
// @Success 200 {object} dto.Token "Access token"
// @Failure 400 {object} dto.JsonError "Incorrect data"
// @Failure 401 {object} dto.JsonError "Incorrect email or password"
// @Failure 404 {object} dto.JsonError "This user wasn't found."
// @Failure 500 {object} dto.JsonError "Something going wrong..."
// @Router /account/auth/login [post]
func (a *Auth) Login(c *gin.Context) {
	ctx := gins.GetCtx(c)
	var body dto.Login

	if err := c.ShouldBindJSON(&body); err != nil {
		gins.Abort(c, badReqErr.WithErr(err))
		return
	}

	if err := validator.Struct(body, validator.Password); err != nil {
		gins.Abort(c, err)
		return
	}

	user := conv.EntityLogin(body)

	tokens, err := a.usecase.Login(ctx, user)
	if err != nil {
		gins.Abort(c, err)
		return
	}

	c.SetCookie(cookieName, tokens.Refresh, cookieAge, cookiePath, cookieHost, false, true)

	result := conv.DtoToken(tokens)

	c.JSON(ok, result)
}

// @Summary Log out a user
// @Description Logs out a user by invalidating the session
// @Tags auth
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} dto.Message "Logout success."
// @Failure 500 {object} dto.JsonError "Something going wrong..."
// @Router  /account/auth/logout [post]
func (a *Auth) Logout(c *gin.Context) {
	c.SetCookie(cookieName, "", -1, cookiePath, cookieHost, false, true)
}

// @Summary Refresh user tokens
// @Description Refreshes the user's tokens using the refresh token from the cookie
// @Tags auth
// @Accept json
// @Produce json
// @Success 200 {object} dto.Token "Refresh token"
// @Failure 401 {object} dto.JsonError "Token is invalid"
// @Failure 404 {object} dto.JsonError "Your token wasn't found., This user wasn't found."
// @Failure 500 {object} dto.JsonError "Something going wrong..."
// @Router /account/auth/refresh [get]
func (a *Auth) Refresh(c *gin.Context) {
	ctx := gins.GetCtx(c)
	refresh, cookieErr := c.Cookie(cookieName)
	if cookieErr != nil {
		gins.Abort(c, badReqErr.WithErr(cookieErr))
		return
	}

	tokens, err := a.usecase.Refresh(ctx, refresh)
	if err != nil {
		gins.Abort(c, err)
		return
	}

	c.SetCookie(cookieName, tokens.Refresh, cookieAge, cookiePath, cookieHost, false, true)

	result := conv.DtoToken(tokens)

	c.JSON(ok, result)
}
