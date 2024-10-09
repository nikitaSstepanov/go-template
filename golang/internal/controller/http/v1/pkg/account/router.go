package account

import "github.com/gin-gonic/gin"

func InitRoutes(h *gin.RouterGroup, a AccountHandler, mid Middleware) *gin.RouterGroup {
	router := h.Group("/account")
	{
		router.GET("/", mid.CheckAccess(), a.Get)
		router.POST("/new", a.Create)
		router.PATCH("/edit", mid.CheckAccess(), a.Update)
		router.GET("/verify/confirm/:code", mid.CheckAccess(), a.Verify)
		router.GET("/verify/resend", mid.CheckAccess(), a.ResendCode)
		router.DELETE("/delete", mid.CheckAccess(), a.Delete)
	}

	return router
}
