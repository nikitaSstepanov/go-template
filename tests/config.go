package tests

import (
	"os"
	"path/filepath"

	"github.com/nikitaSstepanov/tools/client/pg"
	"github.com/nikitaSstepanov/tools/client/redis"
)

type TestConfig struct {
	//Server   hserv.Config        `yaml:"server"`
	Postgres pg.Config    `yaml:"postgres"`
	Redis    redis.Config `yaml:"redis"`
}

func getConfigPath() string {
	path := os.Getenv("CONFIG_PATH_TEST")
	if path == "" {
		path, _ = filepath.Abs("../config/test.yaml")
		return path
	}
	return path
}
