package v1

import (
	"github.com/gin-gonic/gin"
	_ "app/docs"
	swaggerfiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

func (c *Controller) initAccountRoutes(h *gin.RouterGroup) *gin.RouterGroup {
	router := h.Group("/account")
	{
		router.GET("/", c.mid.CheckAccess(), c.account.Get)
		router.PATCH("/edit", c.mid.CheckAccess(), c.account.Update)
		router.POST("/new", c.account.Create)
		router.GET("/verify/confirm/:code", c.mid.CheckAccess(), c.account.Verify)
		router.GET("/verify/resend", c.mid.CheckAccess(), c.account.ResendCode)
		router.DELETE("/delete", c.mid.CheckAccess(), c.account.Delete)
	}

	return router
}

func (c *Controller) initSwaggerRoute(h *gin.RouterGroup) *gin.RouterGroup {
	router := h.Group("swagger")
	{
		router.GET("/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	}

	return router
}

func (c *Controller) initAuthRoutes(h *gin.RouterGroup) *gin.RouterGroup {
	router := h.Group("/auth")
	{
		router.POST("/login", c.auth.Login)
		router.POST("/logout", c.mid.CheckAccess(), c.auth.Logout)
		router.GET("/refresh", c.auth.Refresh)
	}

	return router
}
