package app

import (
	"os"

	"app/internal/controller"
	"app/internal/usecase"
	"app/internal/usecase/storage"

	"github.com/gosuit/confy"
	"github.com/gosuit/httper"
	"github.com/gosuit/sl"
	"github.com/joho/godotenv"
)

type appConfig struct {
	Controller controller.Config `confy:"controller"`
	UseCase    usecase.Config    `confy:"usecase"`
	Storage    storage.Config    `confy:"storage"`
	Server     httper.ServerCfg  `confy:"server"`
	Logger     sl.Config         `confy:"logger"`
}

func getConfig() (*appConfig, error) {
	env := getEnv()

	if env == "LOCAL" {
		if err := godotenv.Load(".env"); err != nil {
			return nil, err
		}
	}

	path := getConfigPath(env)

	var cfg appConfig

	if err := confy.Get(path, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func getEnv() string {
	env, ok := os.LookupEnv("ENVIRONMENT")

	if !ok {
		return "LOCAL"
	}

	return env
}

func getConfigPath(env string) string {
	switch env {

	case "LOCAL":
		return "config/local"

	case "DOCKER":
		return "config/docker"

	default:
		return "config/local"

	}
}
