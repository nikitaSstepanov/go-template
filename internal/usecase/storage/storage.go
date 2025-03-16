package storage

import (
	code "app/internal/usecase/storage/activation_code"
	"app/internal/usecase/storage/token"
	"app/internal/usecase/storage/user"
	"github.com/nikitaSstepanov/tools/client/pg"
	rs "github.com/nikitaSstepanov/tools/client/redis"
	e "github.com/nikitaSstepanov/tools/error"
)

type Storage struct {
	pg     pg.Client
	rs     rs.Client
	Users  *user.User
	Tokens *token.Token
	Codes  *code.Code
}

func New(postgres pg.Client, redis rs.Client) *Storage {
	return &Storage{
		pg:     postgres,
		rs:     redis,
		Users:  user.New(postgres, redis),
		Tokens: token.New(redis),
		Codes:  code.New(redis),
	}
}

func (s *Storage) Close() e.Error {
	s.pg.Close()
	return e.E(s.rs.Close())
}
