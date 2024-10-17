package v1

import (
	"github.com/gin-gonic/gin"
)

func (c *Controller) initAccountRoutes(h *gin.RouterGroup) *gin.RouterGroup {
	router := h.Group("/account")
	{
		router.GET("/", c.mid.CheckAccess(), c.account.Get)
		router.POST("/new", c.account.Create)
		router.PATCH("/edit", c.mid.CheckAccess(), c.account.Update)
		router.GET("/verify/confirm/:code", c.mid.CheckAccess(), c.account.Verify)
		router.GET("/verify/resend", c.mid.CheckAccess(), c.account.ResendCode)
		router.DELETE("/delete", c.mid.CheckAccess(), c.account.Delete)
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
