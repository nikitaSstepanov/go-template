package usecase

import (
	"github.com/nikitaSstepanov/templates/golang/internal/usecase/mail"
	"github.com/nikitaSstepanov/templates/golang/internal/usecase/pkg/account"
	"github.com/nikitaSstepanov/templates/golang/internal/usecase/pkg/auth"
	"github.com/nikitaSstepanov/templates/golang/internal/usecase/storage"
	gomail "github.com/nikitaSstepanov/tools/client/mail"
)

type UseCase struct {
	Account *account.Account
	Auth    *auth.Auth
}

func New(storage *storage.Storage, jwtCfg *auth.JwtOptions, mailCfg *gomail.Config) *UseCase {
	jwt := auth.NewJwt(jwtCfg)
	mail := mail.New(mailCfg)

	return &UseCase{
		Account: account.New(storage.Users, storage.Tokens, storage.Codes, jwt, mail),
		Auth:    auth.New(storage.Users, storage.Tokens, jwt),
	}
}
