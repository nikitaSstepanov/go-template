package auth

import (
	"app/internal/entity"

	"github.com/gosuit/e"
	"github.com/gosuit/lec"
	"github.com/gosuit/utils/coder"
)

type Auth struct {
	user  UserStorage
	jwt   *Jwt
	coder *coder.Coder
}

func New(uc *UseCases, store *Storages) *Auth {
	return &Auth{
		user:  store.User,
		jwt:   uc.Jwt,
		coder: uc.Coder,
	}
}

func (a *Auth) Login(ctx lec.Context, user *entity.User) (*entity.Tokens, e.Error) {
	candidate, err := a.user.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if candidate.Id == 0 {
		return nil, badDataErr
	}

	if err := a.coder.CompareHash(candidate.Password, user.Password); err != nil {
		return nil, badDataErr.WithErr(err)
	}

	access, err := a.jwt.GenerateToken(candidate, accessExpires, false)
	if err != nil {
		return nil, err
	}

	refresh, err := a.jwt.GenerateToken(candidate, refreshExpires, true)
	if err != nil {
		return nil, err
	}

	result := &entity.Tokens{
		Access:  access,
		Refresh: refresh,
	}

	return result, nil
}

func (a *Auth) Refresh(ctx lec.Context, refresh string) (*entity.Tokens, e.Error) {
	claims, err := a.jwt.ValidateToken(refresh, true)
	if err != nil {
		return nil, err
	}

	user, err := a.user.GetById(ctx, claims.Id)
	if err != nil {
		return nil, err
	}

	access, err := a.jwt.GenerateToken(user, accessExpires, false)
	if err != nil {
		return nil, err
	}

	refresh, err = a.jwt.GenerateToken(user, refreshExpires, true)
	if err != nil {
		return nil, err
	}

	result := &entity.Tokens{
		Access:  access,
		Refresh: refresh,
	}

	return result, nil
}

func (a *Auth) ValidateToken(jwtString string, isRefresh bool) (*Claims, e.Error) {
	return a.jwt.ValidateToken(jwtString, isRefresh)
}
