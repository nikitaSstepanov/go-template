package auth

import (
	"app/internal/entity"
	"context"

	"github.com/gosuit/e"
	"golang.org/x/crypto/bcrypt"
)

type Auth struct {
	user  UserStorage
	token TokenStorage
	jwt   *Jwt
}

func New(user UserStorage, token TokenStorage, jwt *Jwt) *Auth {
	return &Auth{
		user:  user,
		token: token,
		jwt:   jwt,
	}
}

func (a *Auth) Login(ctx context.Context, user *entity.User) (*entity.Tokens, e.Error) {
	candidate, err := a.user.GetByEmail(ctx, user.Email)
	if err != nil {
		return nil, err
	}

	if candidate.Id == 0 {
		return nil, badDataErr
	}

	if err := a.checkPassword(candidate.Password, user.Password); err != nil {
		return nil, badDataErr.WithErr(err)
	}

	access, err := a.jwt.GenerateToken(candidate.Id, candidate.Role, accessExpires, false)
	if err != nil {
		return nil, err
	}

	refresh, err := a.jwt.GenerateToken(candidate.Id, candidate.Role, refreshExpires, true)
	if err != nil {
		return nil, err
	}

	token := &entity.Token{
		Token:  refresh,
		UserId: candidate.Id,
	}

	err = a.token.Set(ctx, token)
	if err != nil {
		return nil, err
	}

	result := &entity.Tokens{
		Access:  access,
		Refresh: refresh,
	}

	return result, nil
}

func (a *Auth) Logout(ctx context.Context, userId uint64) e.Error {
	return a.token.Del(ctx, userId)
}

func (a *Auth) Refresh(ctx context.Context, refresh string) (*entity.Tokens, e.Error) {
	claims, err := a.jwt.ValidateToken(refresh, true)
	if err != nil {
		return nil, err
	}

	token, err := a.token.Get(ctx, claims.Id)
	if err != nil {
		return nil, err
	}

	if claims.Id != token.UserId {
		return nil, unauthErr
	}

	user, err := a.user.GetById(ctx, token.UserId)
	if err != nil {
		return nil, err
	}

	access, err := a.jwt.GenerateToken(user.Id, user.Role, accessExpires, false)
	if err != nil {
		return nil, err
	}

	refresh, err = a.jwt.GenerateToken(user.Id, user.Role, refreshExpires, true)
	if err != nil {
		return nil, err
	}

	token = &entity.Token{
		Token:  refresh,
		UserId: user.Id,
	}

	err = a.token.Set(ctx, token)
	if err != nil {
		return nil, err
	}

	result := &entity.Tokens{
		Access:  access,
		Refresh: refresh,
	}

	return result, nil
}

func (a *Auth) checkPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
