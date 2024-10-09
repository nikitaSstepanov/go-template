package auth

import "github.com/gin-gonic/gin"

func InitRoutes(h *gin.RouterGroup, a AuthHandler, mid Middleware) *gin.RouterGroup {
	router := h.Group("/auth")
	{
		router.POST("/login", a.Login)
		router.POST("/logout", mid.CheckAccess(), a.Logout)
		router.GET("/refresh", a.Refresh)
	}

	return router
}
