package app

import (
	"os"

	"app/internal/controller"
	"app/internal/usecase"
	"app/internal/usecase/storage"

	"github.com/gosuit/httper"
	"github.com/gosuit/sl"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

type Config struct {
	Controller controller.Config `yaml:"controller"`
	UseCase    usecase.Config    `yaml:"usecase"`
	Storage    storage.Config    `yaml:"storage"`
	Server     httper.ServerCfg  `yaml:"server"`
	Logger     sl.Config         `yaml:"logger"`
}

func getConfig() (*Config, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	path := getConfigPath()

	var cfg Config

	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

func getConfigPath() string {
	path := os.Getenv("CONFIG_PATH")

	if path == "" {
		return "config/local.yaml"
	}

	return path
}
