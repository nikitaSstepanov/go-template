package token

import (
	"context"

	"github.com/nikitaSstepanov/templates/golang/internal/entity"
	rs "github.com/nikitaSstepanov/tools/client/redis"
	e "github.com/nikitaSstepanov/tools/error"
)

type Token struct {
	redis rs.Client
}

func New(redis rs.Client) *Token {
	return &Token{
		redis,
	}
}

func (t *Token) Get(ctx context.Context, userId uint64) (*entity.Token, e.Error) {
	var token entity.Token

	err := t.redis.Get(ctx, redisKey(userId)).Scan(&token)
	if err != nil {
		if err == rs.Nil {
			return nil, notFoundErr
		} else {
			return nil, internalErr
		}
	}

	return &token, nil
}

func (t *Token) Set(ctx context.Context, token *entity.Token) e.Error {
	err := t.redis.Set(ctx, redisKey(token.UserId), token, redisExpires).Err()
	if err != nil {
		return internalErr
	}

	return nil
}

func (t *Token) Del(ctx context.Context, userId uint64) e.Error {
	err := t.redis.Del(ctx, redisKey(userId)).Err()
	if err != nil {
		return internalErr
	}

	return nil
}
