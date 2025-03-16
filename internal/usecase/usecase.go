package usecase

import (
	"app/internal/usecase/mail"
	"app/internal/usecase/pkg/account"
	"app/internal/usecase/pkg/auth"
	"app/internal/usecase/storage"

	gomail "github.com/gosuit/mail"
)

type UseCase struct {
	Account *account.Account
	Auth    *auth.Auth
}

type Config struct {
	Jwt  auth.JwtOptions `yaml:"jwt"`
	Mail gomail.Config   `yaml:"mail"`
}

func New(storage *storage.Storage, cfg *Config) *UseCase {
	jwt := auth.NewJwt(&cfg.Jwt)
	mail := mail.New(&cfg.Mail)

	return &UseCase{
		Account: account.New(storage.Users, storage.Tokens, storage.Codes, jwt, mail),
		Auth:    auth.New(storage.Users, storage.Tokens, jwt),
	}
}
