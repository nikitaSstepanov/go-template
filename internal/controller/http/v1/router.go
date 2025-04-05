package v1

import (
	"app/internal/controller/http/v1/middleware"
	"app/internal/controller/http/v1/pkg/account"
	"app/internal/controller/http/v1/pkg/auth"
	"app/internal/entity/types"
	"app/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/gosuit/gins"
	"github.com/gosuit/httper"
	"github.com/gosuit/lec"
)

type Router struct {
	account AccountHandler
	auth    AuthHandler
	mid     Middleware
}

type Config struct {
	Swagger       SwaggerCfg    `confy:"swagger"`
	RefreshCookie httper.Cookie `confy:"refresh_cookie"`
}

func New(uc *usecase.UseCase, cfg *Config) *Router {
	setSwaggerConfig(cfg.Swagger)

	return &Router{
		account: account.New(uc.User, &cfg.RefreshCookie),
		auth:    auth.New(uc.Auth, &cfg.RefreshCookie),
		mid:     middleware.New(uc.Auth),
	}
}

func (r *Router) InitRoutes(c lec.Context, h *gin.RouterGroup) *gin.RouterGroup {
	router := h.Group("/v1")
	{
		router.Use(gins.InitLogger(c))

		account := r.initAccountRoutes(router)
		r.initAuthRoutes(account)

		r.initSwaggerRoute(router)
	}

	return router
}

func (r *Router) initAccountRoutes(h *gin.RouterGroup) *gin.RouterGroup {
	router := h.Group("/account")
	{
		router.PUT("/edit", r.mid.CheckAccess(), r.account.Update)
		router.PATCH("/edit/role", r.mid.CheckAccess(types.ADMIN), r.account.SetRole)
		router.POST("/new", r.account.Create)
		router.GET("/verify/confirm/:code", r.mid.CheckAccess(), r.account.Verify)
		router.GET("/verify/resend", r.mid.CheckAccess(), r.account.ResendCode)
		router.DELETE("/delete", r.mid.CheckAccess(), r.account.Delete)
	}

	h.GET("/account", r.mid.CheckAccess(), r.account.Get)

	return router
}

func (r *Router) initAuthRoutes(h *gin.RouterGroup) *gin.RouterGroup {
	router := h.Group("/auth")
	{
		router.POST("/login", r.auth.Login)
		router.POST("/logout", r.auth.Logout)
		router.GET("/refresh", r.auth.Refresh)
	}

	return router
}
