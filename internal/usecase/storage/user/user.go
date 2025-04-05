package user

import (
	"app/internal/entity"
	"app/internal/entity/types"

	"github.com/gosuit/e"
	"github.com/gosuit/lec"
	"github.com/gosuit/pg"
	"github.com/gosuit/rs"
)

type User struct {
	postgres pg.Client
	redis    rs.Client
}

func New(postgres pg.Client, redis rs.Client) *User {
	return &User{
		postgres,
		redis,
	}
}

func (u *User) GetById(ctx lec.Context, id uint64) (*entity.User, e.Error) {
	var user entity.User

	err := u.redis.Get(ctx, redisKey(id)).Scan(&user)
	if err != nil && err != rs.Nil {
		return nil, e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}

	if user.Id != 0 {
		return &user, nil
	}

	query, args := idQuery(id)

	row := u.postgres.QueryRow(ctx, query, args...)

	if err := user.Scan(row); err != nil {
		if err == pg.ErrNoRows {
			return nil, notFoundErr.
				WithCtx(ctx).
				WithErr(err)
		} else {
			return nil, e.InternalErr.
				WithCtx(ctx).
				WithErr(err)
		}
	}

	err = u.redis.Set(ctx, redisKey(id), &user, redisExpires).Err()
	if err != nil {
		return nil, e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}

	return &user, nil
}

func (u *User) GetByEmail(ctx lec.Context, email string) (*entity.User, e.Error) {
	var user entity.User

	query, args := emailQuery(email)

	row := u.postgres.QueryRow(ctx, query, args...)

	if err := user.Scan(row); err != nil {
		if err == pg.ErrNoRows {
			return nil, notFoundErr.WithErr(err)
		}

		return nil, e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}

	return &user, nil
}

func (u *User) Create(ctx lec.Context, user *entity.User) e.Error {
	query, args := createQuery(user)

	tx, err := u.postgres.Begin(ctx)
	if err != nil {
		return e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}
	defer tx.Rollback(ctx)

	row := tx.QueryRow(ctx, query, args...)

	err = row.Scan(&user.Id)
	if err != nil {
		return e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}

	if err := tx.Commit(ctx); err != nil {
		return e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}

	user.Role = types.USER

	err = u.redis.Set(ctx, redisKey(user.Id), user, redisExpires).Err()
	if err != nil {
		return e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}

	return nil
}

func (u *User) Update(ctx lec.Context, user *entity.User) e.Error {
	query, args := updateQuery(user)

	tx, err := u.postgres.Begin(ctx)
	if err != nil {
		return e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}
	defer tx.Rollback(ctx)

	if _, err = tx.Exec(ctx, query, args...); err != nil {
		return e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}

	if err := tx.Commit(ctx); err != nil {
		return e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}

	err = u.redis.Del(ctx, redisKey(user.Id)).Err()
	if err != nil {
		return e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}

	err = u.redis.Set(ctx, redisKey(user.Id), user, redisExpires).Err()
	if err != nil {
		return e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}

	return nil
}

func (u *User) Delete(ctx lec.Context, id uint64) e.Error {
	query, args := deleteQuery(id)

	tx, err := u.postgres.Begin(ctx)
	if err != nil {
		return e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, query, args...)
	if err != nil {
		return e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}

	if err = tx.Commit(ctx); err != nil {
		return e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}

	err = u.redis.Del(ctx, redisKey(id)).Err()
	if err != nil {
		return e.InternalErr.
			WithCtx(ctx).
			WithErr(err)
	}

	return nil
}
