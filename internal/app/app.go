package app

import (
	"fmt"

	"app/internal/controller"
	"app/internal/usecase"
	"app/internal/usecase/storage"

	"github.com/gosuit/e"
	"github.com/gosuit/httper"
	"github.com/gosuit/lec"
	"github.com/gosuit/sl"
)

type App struct {
	controller *controller.Controller
	usecase    *usecase.UseCase
	storage    *storage.Storage
	server     *httper.Server
	ctx        lec.Context
}

func New() *App {
	cfg, err := getConfig()
	if err != nil {
		panic(fmt.Sprintf("Can`t get app config. Error: %v", err))
	}

	log := sl.New(&cfg.Logger)

	ctx := lec.New(log)

	app := &App{}

	app.ctx = ctx

	app.storage = storage.New(ctx, &cfg.Storage)

	app.usecase = usecase.New(app.storage, &cfg.UseCase)

	app.controller = controller.New(app.usecase, &cfg.Controller)

	handler := app.controller.InitRoutes(ctx)

	app.server = httper.NewServer(&cfg.Server, handler)

	return app
}

func (a *App) Run() {
	log := a.ctx.Logger()

	a.server.Start()

	if err := a.shutdown(); err != nil {
		log.Error("Failed to shutdown app", sl.ErrAttr(err))
		return
	}

	log.Info("Application stopped successfully")
}

func (a *App) shutdown() e.Error {
	log := a.ctx.Logger()

	err := e.E(a.server.Shutdown(log.ToSlog()))
	if err != nil {
		log.Error("Failed to stop http server", err.SlErr())
		return err
	}

	if err := a.storage.Close(); err != nil {
		log.Error("Failed to close storage", err.SlErr())
		return err
	}

	return nil
}
