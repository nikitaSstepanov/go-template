package v1

import (
	_ "app/docs"
	"app/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/gosuit/gins"
	"github.com/gosuit/lec"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Router struct {
	account AccountHandler
	auth    AuthHandler
	mid     Middleware
}

type Config struct{}

func New(uc *usecase.UseCase, cfg *Config) *Router {
	return &Router{}
}

func (r *Router) InitRoutes(c lec.Context, h *gin.RouterGroup) *gin.RouterGroup {
	router := h.Group("/v1")
	{
		router.Use(gins.InitLogger(c))

		r.initSwaggerRoute(router)
		r.initAccountRoutes(router)
		r.initAuthRoutes(router)
	}

	return router
}

func (r *Router) initAccountRoutes(h *gin.RouterGroup) *gin.RouterGroup {
	router := h.Group("/account")
	{
		router.GET("/", r.mid.CheckAccess(), r.account.Get)
		router.PATCH("/edit", r.mid.CheckAccess(), r.account.Update)
		router.POST("/new", r.account.Create)
		router.GET("/verify/confirm/:code", r.mid.CheckAccess(), r.account.Verify)
		router.GET("/verify/resend", r.mid.CheckAccess(), r.account.ResendCode)
		router.DELETE("/delete", r.mid.CheckAccess(), r.account.Delete)
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

func (r *Router) initAuthRoutes(h *gin.RouterGroup) *gin.RouterGroup {
	router := h.Group("/auth")
	{
		router.POST("/login", r.auth.Login)
		router.POST("/logout", r.mid.CheckAccess(), r.auth.Logout)
		router.GET("/refresh", r.auth.Refresh)
	}

	return router
}
