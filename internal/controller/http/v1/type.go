package v1

import (
	"app/internal/entity/types"

	"github.com/gin-gonic/gin"
)

type AccountHandler interface {
	Get(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	SetRole(ctx *gin.Context)
	Verify(ctx *gin.Context)
	ResendCode(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type AuthHandler interface {
	Login(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Refresh(ctx *gin.Context)
}

type Middleware interface {
	CheckAccess(roles ...types.Role) gin.HandlerFunc
}
