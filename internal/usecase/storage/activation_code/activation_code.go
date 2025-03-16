package activation_code

import (
	"app/internal/entity"

	"github.com/gosuit/e"
	"github.com/gosuit/lec"
	"github.com/gosuit/rs"
)

type Code struct {
	redis rs.Client
}

func New(redis rs.Client) *Code {
	return &Code{
		redis,
	}
}

func (c *Code) Get(ctx lec.Context, userId uint64) (*entity.ActivationCode, e.Error) {
	var result entity.ActivationCode

	err := c.redis.Get(ctx, redisKey(userId)).Scan(&result)
	if err != nil {
		if err == rs.Nil {
			return nil, notFoundErr.
				WithCtx(ctx).
				WithErr(err)
		} else {
			return nil, e.InternalErr.
				WithCtx(ctx).
				WithErr(err)
		}
	}

	return &result, nil
}

func (c *Code) Set(ctx lec.Context, code *entity.ActivationCode) e.Error {
	err := c.redis.Set(ctx, redisKey(code.UserId), code, redisExpires).Err()
	if err != nil {
		return e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}

	return nil
}

func (c *Code) Del(ctx lec.Context, userId uint64) e.Error {
	err := c.redis.Del(ctx, redisKey(userId)).Err()
	if err != nil {
		return e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}

	return nil
}
