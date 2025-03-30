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

type Config struct {
	Controller controller.Config `confy:"controller"`
	UseCase    usecase.Config    `confy:"usecase"`
	Storage    storage.Config    `confy:"storage"`
	Server     httper.ServerCfg  `confy:"server"`
	Logger     sl.Config         `confy:"logger"`
}

func getConfig() (*Config, error) {
	env := os.Getenv("ENVIRONMENT")

	if env == "LOCAL" {
		if err := godotenv.Load(".env"); err != nil {
			return nil, err
		}
	}

	path := getConfigPath(env)

	var cfg Config

	if err := confy.Get(path, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func getConfigPath(env string) string {
	switch env {

	case "LOCAL":
		return "config/local"

	case "DOCKER":
		return "config/local"

	default:
		return "config/local"

	}
}
