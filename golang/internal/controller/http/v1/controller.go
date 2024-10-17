package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/middleware"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/pkg/account"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/pkg/auth"
	"github.com/nikitaSstepanov/templates/golang/internal/usecase"
	jwt "github.com/nikitaSstepanov/templates/golang/internal/usecase/pkg/auth"
)

type Controller struct {
	account AccountHandler
	auth    AuthHandler
	mid     Middleware
}

func New(ctx context.Context, uc *usecase.UseCase, jwtOpts *jwt.JwtOptions) *Controller {
	return &Controller{
		account: account.New(uc.Account),
		auth:    auth.New(uc.Auth),
		mid:     middleware.New(jwtOpts),
	}
}

func (c *Controller) InitRoutes(ctx context.Context, mode string) *gin.Engine {
	setGinMode(mode)

	router := gin.New()

	if gin.Mode() != gin.ReleaseMode {
		router.Use(c.mid.InitLogger(ctx))
	}

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	api := router.Group("/api/v1")
	{
		account := c.initAccountRoutes(api)

		c.initAuthRoutes(account)
	}

	return router
}

func setGinMode(mode string) {
	switch mode {

	case "RELEASE":
		gin.SetMode(gin.ReleaseMode)

	case "TEST":
		gin.SetMode(gin.TestMode)

	case "DEBUG":
		gin.SetMode(gin.DebugMode)

	default:
		gin.SetMode(gin.DebugMode)

	}
}
