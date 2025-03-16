package v1

import (
	"github.com/gin-gonic/gin"
)

type AuthHandler interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Refresh(ctx *gin.Context)
}

type AccountHandler interface {
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Verify(ctx *gin.Context)
	ResendCode(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type Middleware interface {
	CheckAccess() gin.HandlerFunc
}
