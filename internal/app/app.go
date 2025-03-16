package app

import (
	"context"
	"fmt"

	controller "app/internal/controller/http/v1"
	"app/internal/usecase"
	"app/internal/usecase/storage"
	"app/pkg/swagger"
	"github.com/nikitaSstepanov/tools/client/pg"
	rs "github.com/nikitaSstepanov/tools/client/redis"
	e "github.com/nikitaSstepanov/tools/error"
	"github.com/nikitaSstepanov/tools/hserv"
	"github.com/nikitaSstepanov/tools/migrate"
	"github.com/nikitaSstepanov/tools/sl"
)

type App struct {
	controller *controller.Controller
	usecase    *usecase.UseCase
	storage    *storage.Storage
	server     *hserv.Server
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
		logger.Error("Can`t connect to postgres", sl.ErrAttr(err))
	} else {
		logger.Info("Connect to postgres succesfully")
	}

	if err := migrate.MigratePg(pg, "./migrations"); err != nil {
		logger.Error("Can`t migrate postgres scheme.", sl.ErrAttr(err))
	} else {
		logger.Info("Postgres scheme migrated")
	}

	redis, err := rs.ConnectToRedis(ctx, &cfg.Redis)
	if err != nil {
		logger.Error("Can`t connect to redis.", sl.ErrAttr(err))
	} else {
		logger.Info("Connect to redis succesfully")
	}

	swagger.SetSwaggerConfig(cfg.Swagger)

	app := &App{}

	app.ctx = ctx

	app.storage = storage.New(pg, redis)

	app.usecase = usecase.New(app.storage, &cfg.Jwt, &cfg.Mail)

	app.controller = controller.New(ctx, app.usecase, &cfg.Jwt)

	handler := app.controller.InitRoutes(ctx, cfg.Mode)

	app.server = hserv.New(handler, &cfg.Server)

	return app
}

func (a *App) Run() {
	log := sl.L(a.ctx)

	a.server.Start()

	if err := a.shutdown(); err != nil {
		log.Error("Failed to shutdown app", sl.ErrAttr(err))
		return
	}

	log.Info("Application stopped successfully")
}

func (a *App) shutdown() e.Error {
	log := sl.L(a.ctx)

	err := e.E(a.server.Shutdown(a.ctx))
	if err != nil {
		log.Error("Failed to stop http server", err.SlErr())
		return err
	}

	err = a.storage.Close()
	if err != nil {
		log.Error("Failed to close storage", err.SlErr())
		return err
	}

	return nil
}
