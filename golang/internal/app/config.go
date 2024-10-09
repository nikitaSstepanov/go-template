package app

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/nikitaSstepanov/templates/golang/internal/usecase/pkg/auth"
	"github.com/nikitaSstepanov/tools/client/mail"
	"github.com/nikitaSstepanov/tools/client/pg"
	"github.com/nikitaSstepanov/tools/client/redis"
	server "github.com/nikitaSstepanov/tools/http_server"
	"github.com/nikitaSstepanov/tools/log"
)

type appConfig struct {
	Server   server.Config   `yaml:"server"`
	Postgres pg.Config       `yaml:"postgres"`
	Redis    redis.Config    `yaml:"redis"`
	Mail     mail.Config     `yaml:"mail"`
	Jwt      auth.JwtOptions `yaml:"jwt"`
	Logger   log.Config      `yaml:"logger"`
	Mode     string          `yaml:"mode" env:"MODE" env-default:"DEBUG"`
}

func getAppConfig() (*appConfig, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, err
	}

	path := getConfigPath()

	var cfg appConfig

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
