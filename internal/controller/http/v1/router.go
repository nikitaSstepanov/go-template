package v1

import (
	"app/internal/controller/http/v1/middleware"
	"app/internal/controller/http/v1/pkg/account"
	"app/internal/controller/http/v1/pkg/auth"
	"app/internal/usecase"
	"app/pkg/swagger"

	"github.com/gin-gonic/gin"
	"github.com/gosuit/gins"
	"github.com/gosuit/httper"
	"github.com/gosuit/lec"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Router struct {
	account AccountHandler
	auth    AuthHandler
	mid     Middleware
}

type Config struct {
	Swagger       swagger.SwaggerSpec `confy:"swagger"`
	RefreshCookie httper.Cookie       `confy:"refresh_cookie"`
}

func New(uc *usecase.UseCase, cfg *Config) *Router {
	swagger.SetSwaggerConfig(cfg.Swagger)

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
		router.PATCH("/edit", r.mid.CheckAccess(), r.account.Update)
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

func (r *Router) initSwaggerRoute(h *gin.RouterGroup) *gin.RouterGroup {
	router := h.Group("swagger")
	{
		router.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	return router
}
