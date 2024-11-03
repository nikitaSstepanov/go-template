package tests

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/nikitaSstepanov/tools/client/pg"
	"github.com/nikitaSstepanov/tools/client/redis"
)

var testCfg = newTestConf()

type TestConfig struct {
	Server   Server       `yaml:"server"`
	Postgres pg.Config    `yaml:"postgres"`
	Redis    redis.Config `yaml:"redis"`
}

func (cfg *TestConfig) ToURL() string {
	u := url.URL{
		Scheme: cfg.Server.Scheme,
		Host:   fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port),
		Path:   cfg.Server.Path,
	}

	return u.String()
}

func newTestConf() TestConfig {
	var cfg = TestConfig{}
	path := getConfigPath()

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		panic(fmt.Sprintf("Failed to read test config: %s", err.Error()))
	}

	return cfg
}

type Server struct {
	Scheme string `yaml:"scheme"`
	Host   string `yaml:"host"`
	Port   string `yaml:"port"`
	Path   string `yaml:"path"`
}

func getConfigPath() string {
	path := os.Getenv("CONFIG_PATH_TEST")
	if path == "" {
		path, _ = filepath.Abs("../config/test.yaml")
		return path
	}
	return path
}
