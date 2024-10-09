package app

import (
	"context"
	"fmt"
	"log/slog"

	controller "github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1"
	"github.com/nikitaSstepanov/templates/golang/internal/usecase"
	"github.com/nikitaSstepanov/templates/golang/internal/usecase/storage"
	"github.com/nikitaSstepanov/tools/client/pg"
	rs "github.com/nikitaSstepanov/tools/client/redis"
	server "github.com/nikitaSstepanov/tools/http_server"
	"github.com/nikitaSstepanov/tools/log"
	"github.com/nikitaSstepanov/tools/migrate"
)

type App struct {
	controller *controller.Controller
	usecase    *usecase.UseCase
	storage    *storage.Storage
	server     *server.Server
	logger     *slog.Logger
}

func New() *App {
	cfg, err := getAppConfig()
	if err != nil {
		panic(fmt.Errorf("can`t get application config. Error: %s", err.Error()))
	}

	logger := log.New(&cfg.Logger)

	ctx := context.TODO()

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

	app.logger = logger

	app.storage = storage.New(pg, redis)

	app.usecase = usecase.New(app.storage, &cfg.Jwt, &cfg.Mail)

	app.controller = controller.New(app.usecase, logger, &cfg.Jwt)

	handler := app.controller.InitRoutes(cfg.Mode)

	app.server = server.New(handler, &cfg.Server)

	return app
}

func (a *App) Run() {
	a.server.Start()

	a.shutdown()
}

func (a *App) shutdown() error {
	return a.server.Shutdown(a.logger)
}
