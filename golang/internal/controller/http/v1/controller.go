package v1

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/middleware"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/pkg/account"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/pkg/auth"
	"github.com/nikitaSstepanov/templates/golang/internal/usecase"
	jwt "github.com/nikitaSstepanov/templates/golang/internal/usecase/pkg/auth"
)

type Controller struct {
	ctx     context.Context
	account *account.Account
	auth    *auth.Auth
	mid     *middleware.Middleware
	log     *slog.Logger
}

func New(ctx context.Context, uc *usecase.UseCase, logger *slog.Logger, jwtOpts *jwt.JwtOptions) *Controller {
	return &Controller{
		ctx:     ctx,
		account: account.New(uc.Account),
		auth:    auth.New(uc.Auth),
		mid:     middleware.New(jwtOpts),
		log:     logger,
	}
}

func (c *Controller) InitRoutes(mode string) *gin.Engine {
	setGinMode(mode)

	router := gin.New()

	if gin.Mode() != gin.ReleaseMode {
		router.Use(c.mid.InitLogger(c.ctx))
	}

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "pong")
	})

	api := router.Group("/api/v1")
	{
		account := account.InitRoutes(api, c.account, c.mid)

		auth.InitRoutes(account, c.auth, c.mid)
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
