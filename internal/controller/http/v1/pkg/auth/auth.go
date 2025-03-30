package auth

import (
	conv "app/internal/controller/http/v1/converter"
	"app/internal/controller/http/v1/dto"
	"app/internal/controller/http/v1/validator"

	"github.com/gin-gonic/gin"
	"github.com/gosuit/gins"
	"github.com/gosuit/httper"
)

type Auth struct {
	usecase AuthUseCase
	cookie  *httper.Cookie
}

func New(uc AuthUseCase, cookie *httper.Cookie) *Auth {
	return &Auth{
		usecase: uc,
		cookie:  cookie,
	}
}

// @Summary Log in a user
// @Description Logs in a user with email and password
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body dto.Login true "Login information"
// @Success 200 {object} dto.Token "Access token"
// @Failure 400 {object} resp.JsonError "Incorrect data"
// @Failure 401 {object} resp.JsonError "Incorrect email or password"
// @Failure 404 {object} resp.JsonError "This user wasn't found."
// @Failure 500 {object} resp.JsonError "Something going wrong..."
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

	c.SetCookie(
		a.cookie.Name, tokens.Refresh, a.cookie.Age, a.cookie.Path,
		a.cookie.Host, a.cookie.Secure, a.cookie.HttpOnly,
	)

	result := conv.DtoToken(tokens)

	c.JSON(ok, result)
}

// @Summary Log out a user
// @Description Logs out a user by invalidating the session
// @Tags Auth
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} resp.Message "Logout success."
// @Failure 500 {object} resp.JsonError "Something going wrong..."
// @Router  /account/auth/logout [post]
func (a *Auth) Logout(c *gin.Context) {
	c.SetCookie(
		a.cookie.Name, "", -1, a.cookie.Path,
		a.cookie.Host, a.cookie.Secure, a.cookie.HttpOnly,
	)
}

// @Summary Refresh user tokens
// @Description Refreshes the user's tokens using the refresh token from the cookie
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} dto.Token "Refresh token"
// @Failure 401 {object} resp.JsonError "Token is invalid"
// @Failure 404 {object} resp.JsonError "Your token wasn't found., This user wasn't found."
// @Failure 500 {object} resp.JsonError "Something going wrong..."
// @Router /account/auth/refresh [get]
func (a *Auth) Refresh(c *gin.Context) {
	ctx := gins.GetCtx(c)

	refresh, cookieErr := c.Cookie(a.cookie.Name)
	if cookieErr != nil {
		gins.Abort(c, badReqErr.WithErr(cookieErr))
		return
	}

	tokens, err := a.usecase.Refresh(ctx, refresh)
	if err != nil {
		gins.Abort(c, err)
		return
	}

	c.SetCookie(
		a.cookie.Name, tokens.Refresh, a.cookie.Age, a.cookie.Path,
		a.cookie.Host, a.cookie.Secure, a.cookie.HttpOnly,
	)

	result := conv.DtoToken(tokens)

	c.JSON(ok, result)
}
