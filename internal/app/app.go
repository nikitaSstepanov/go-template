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

func Run() {
	cfg, err := getConfig()
	if err != nil {
		panic(fmt.Sprintf("Can`t get app config. Error: %v", err))
	}

	log := sl.New(&cfg.Logger)

	ctx := lec.New(log)

	storage := storage.New(ctx, &cfg.Storage)

	usecase := usecase.New(storage, &cfg.UseCase)

	controller := controller.New(usecase, &cfg.Controller)

	handler := controller.InitRoutes(ctx)

	server := httper.NewServer(&cfg.Server, handler)

	server.Start()

	log.Info("Application started successfully")

	if err := e.E(server.Shutdown(log.ToSlog())); err != nil {
		log.Fatal("Failed to stop http server", err.SlErr())
	}

	if err := storage.Close(); err != nil {
		log.Fatal("Failed to close storage", err.SlErr())
	}

	log.Info("Application stopped successfully")
}
