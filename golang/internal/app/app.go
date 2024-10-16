package app

import (
	"context"
	"fmt"

	controller "github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1"
	"github.com/nikitaSstepanov/templates/golang/internal/usecase"
	"github.com/nikitaSstepanov/templates/golang/internal/usecase/storage"
	"github.com/nikitaSstepanov/tools/client/pg"
	rs "github.com/nikitaSstepanov/tools/client/redis"
	server "github.com/nikitaSstepanov/tools/http_server"
	"github.com/nikitaSstepanov/tools/migrate"
	"github.com/nikitaSstepanov/tools/sl"
)

type App struct {
	controller *controller.Controller
	usecase    *usecase.UseCase
	storage    *storage.Storage
	server     *server.Server
	ctx        context.Context
}

func New() *App {
	cfg, err := getAppConfig()
	if err != nil {
		panic(fmt.Errorf("can`t get application config. Error: %s", err.Error()))
	}

	logger := sl.New(&cfg.Logger)

	ctx := sl.ContextWithLogger(context.TODO(), logger)

	pg, err := pg.ConnectToDb(ctx, &cfg.Postgres)
	if err != nil {
		logger.Error("Can`t connect to postgres. Error: " + err.Error())
	} else {
		logger.Info("Connect to postgres succesfully")
	}

	if err := migrate.MigratePg(pg, "./migrations"); err != nil {
		logger.Error("Can`t migrate postgres scheme. Error: " + err.Error())
	} else {
		logger.Info("Postgres scheme migrated")
	}

	redis, err := rs.ConnectToRedis(ctx, &cfg.Redis)
	if err != nil {
		logger.Error("Can`t connect to redis. Error: " + err.Error())
	} else {
		logger.Info("Connect to redis succesfully")
	}

	app := &App{}

	app.ctx = ctx

	app.storage = storage.New(pg, redis)

	app.usecase = usecase.New(app.storage, &cfg.Jwt, &cfg.Mail)

	app.controller = controller.New(ctx, app.usecase, logger, &cfg.Jwt)

	logger.Debug("Try to get server")
	handler := app.controller.InitRoutes(cfg.Mode)


	app.server = server.New(handler, &cfg.Server)

	return app
}

func (a *App) Run() {
	a.server.Start()
	sl.L(a.ctx).Debug("Server start on port")

	a.shutdown()
}

func (a *App) shutdown() error {
	return a.server.Shutdown(a.ctx)
}
