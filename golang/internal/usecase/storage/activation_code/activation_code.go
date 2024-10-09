package activation_code

import (
	"context"

	"github.com/nikitaSstepanov/templates/golang/internal/entity"
	rs "github.com/nikitaSstepanov/tools/client/redis"
	e "github.com/nikitaSstepanov/tools/error"
)

type Code struct {
	redis rs.Client
}

func New(redis rs.Client) *Code {
	return &Code{
		redis,
	}
}

func (c *Code) Get(ctx context.Context, userId uint64) (*entity.ActivationCode, e.Error) {
	var result entity.ActivationCode

	err := c.redis.Get(ctx, redisKey(userId)).Scan(&result)
	if err != nil {
		if err == rs.Nil {
			return nil, notFoundErr
		} else {
			return nil, internalErr
		}
	}

	return &result, nil
}

func (c *Code) Set(ctx context.Context, code *entity.ActivationCode) e.Error {
	err := c.redis.Set(ctx, redisKey(code.UserId), code, redisExpires).Err()
	if err != nil {
		return internalErr
	}

	return nil
}

func (c *Code) Del(ctx context.Context, userId uint64) e.Error {
	err := c.redis.Del(ctx, redisKey(userId)).Err()
	if err != nil {
		return internalErr
	}

	return nil
}
