package controller

import (
	v1 "app/internal/controller/http/v1"
	"app/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/gosuit/httper"
	"github.com/gosuit/lec"
)

type Controller struct {
	v1  *v1.Router
	cfg *Config
}

type Config struct {
	V1   v1.Config `confy:"v1"`
	Mode string    `confy:"mode" env:"MODE" env-default:"DEBUG"`
}

func New(uc *usecase.UseCase, cfg *Config) *Controller {
	return &Controller{
		v1:  v1.New(uc, &cfg.V1),
		cfg: cfg,
	}
}

func (c *Controller) InitRoutes(ctx lec.Context) *gin.Engine {
	setGinMode(c.cfg.Mode)

	router := gin.New()

	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(httper.StatusOK, "pong")
	})

	api := router.Group("/api")
	{
		c.v1.InitRoutes(ctx, api)
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
