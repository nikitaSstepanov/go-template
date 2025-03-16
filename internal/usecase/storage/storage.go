package storage

import (
	code "app/internal/usecase/storage/activation_code"
	"app/internal/usecase/storage/user"

	"github.com/gosuit/e"
	"github.com/gosuit/lec"
	"github.com/gosuit/pg"
	"github.com/gosuit/rs"
	"github.com/gosuit/sl"
)

type Storage struct {
	Users *user.User
	Codes *code.Code
	pg    pg.Client
	rs    rs.Client
}

type Config struct {
	Postgres pg.Config `yaml:"postgres"`
	Redis    rs.Config `yaml:"redis"`
}

func New(c lec.Context, cfg *Config) *Storage {
	postgres := connectPg(c, cfg.Postgres)
	redis := connectRs(c, cfg.Redis)

	return &Storage{
		Users: user.New(postgres, redis),
		Codes: code.New(redis),
		pg:    postgres,
		rs:    redis,
	}
}

func connectPg(c lec.Context, cfg pg.Config) pg.Client {
	log := c.Logger()

	postgres, err := pg.New(c, &cfg)
	if err != nil {
		log.Fatal("Can`t connect to postgres.", sl.ErrAttr(err))
	} else {
		log.Info("Postgres is connected.")
	}

	return postgres
}

func connectRs(c lec.Context, cfg rs.Config) rs.Client {
	log := c.Logger()

	redis, err := rs.New(c, &cfg)
	if err != nil {
		log.Fatal("Can`t connect to redis.", sl.ErrAttr(err))
	} else {
		log.Info("Redis is connected.")
	}

	return redis
}

func (s *Storage) Close() e.Error {
	s.pg.Close()
	return e.E(s.rs.Close())
}
