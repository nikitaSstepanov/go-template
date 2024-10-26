package app

import (
	"os"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
	"github.com/nikitaSstepanov/templates/golang/docs"
	"github.com/nikitaSstepanov/templates/golang/internal/usecase/pkg/auth"
	"github.com/nikitaSstepanov/tools/client/mail"
	"github.com/nikitaSstepanov/tools/client/pg"
	"github.com/nikitaSstepanov/tools/client/redis"
	"github.com/nikitaSstepanov/tools/hserv"
	"github.com/nikitaSstepanov/tools/sl"
)

type appConfig struct {
	Server   hserv.Config    `yaml:"server"`
	Postgres pg.Config       `yaml:"postgres"`
	Redis    redis.Config    `yaml:"redis"`
	Mail     mail.Config     `yaml:"mail"`
	Jwt      auth.JwtOptions `yaml:"jwt"`
	Logger   sl.Config       `yaml:"logger"`
	Mode     string          `yaml:"mode" env:"MODE" env-default:"DEBUG"`
	Swagger  SwaggerSpec     `yaml:"swagger"`
}

type SwaggerSpec struct {
	Version     string   `yaml:"version"`
	Host        string   `yaml:"host"`
	BasePath    string   `yaml:"base_path"`
	Schemes     []string `yaml:"schemes"`
	Title       string   `yaml:"title"`
	Description string   `yaml:"description"`
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

func setSwaggerConfig(cfg SwaggerSpec) {
	info := docs.SwaggerInfo

	info.Version = cfg.Version
	info.Host = cfg.Host
	info.BasePath = cfg.BasePath
	info.Schemes = cfg.Schemes
	info.Title = cfg.Title
	info.Description = cfg.Description
}
