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

	a.shutdown()

	log.Info("Application stopped successfully")
}

func (a *App) shutdown() {
	log := a.ctx.Logger()

	err := e.E(a.server.Shutdown(log.ToSlog()))
	if err != nil {
		log.Fatal("Failed to stop http server", err.SlErr())
	}

	if err := a.storage.Close(); err != nil {
		log.Fatal("Failed to close storage", err.SlErr())
	}
}
