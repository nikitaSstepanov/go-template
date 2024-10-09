package storage

import (
	code "github.com/nikitaSstepanov/templates/golang/internal/usecase/storage/activation_code"
	"github.com/nikitaSstepanov/templates/golang/internal/usecase/storage/token"
	"github.com/nikitaSstepanov/templates/golang/internal/usecase/storage/user"
	"github.com/nikitaSstepanov/tools/client/pg"
	rs "github.com/nikitaSstepanov/tools/client/redis"
)

type Storage struct {
	Users  *user.User
	Tokens *token.Token
	Codes  *code.Code
}

func New(postgres pg.Client, redis rs.Client) *Storage {
	return &Storage{
		Users:  user.New(postgres, redis),
		Tokens: token.New(redis),
		Codes:  code.New(redis),
	}
}
