package usecase

import (
	"app/internal/usecase/mail"
	"app/internal/usecase/pkg/auth"
	"app/internal/usecase/pkg/user"
	"app/internal/usecase/storage"

	gomail "github.com/gosuit/mail"
	"github.com/gosuit/sl"
	"github.com/gosuit/utils/coder"
)

type UseCase struct {
	User *user.User
	Auth *auth.Auth
}

type Config struct {
	Jwt   auth.JwtOptions `confy:"jwt"`
	Mail  gomail.Config   `confy:"mail"`
	Coder coder.Config   
}

func New(storage *storage.Storage, cfg *Config) *UseCase {
	return &UseCase{
		User: setupAccount(storage, cfg),
		Auth: setupAuth(storage, cfg),
	}
}

func setupAccount(storage *storage.Storage, cfg *Config) *user.User {
	jwt := auth.NewJwt(&cfg.Jwt)
	mail := mail.New(&cfg.Mail)

	coder, err := coder.New(&cfg.Coder)
	if err != nil {
		sl.Default().Fatal("Can`t init coder.", sl.ErrAttr(err))
	}

	return user.New(
		&user.UseCases{
			Jwt:   jwt,
			Mail:  mail,
			Coder: coder,
		},
		&user.Storages{
			User: storage.Users,
			Code: storage.Codes,
		},
	)
}

func setupAuth(storage *storage.Storage, cfg *Config) *auth.Auth {
	jwt := auth.NewJwt(&cfg.Jwt)

	coder, err := coder.New(&cfg.Coder)
	if err != nil {
		sl.Default().Fatal("Can`t init coder.", sl.ErrAttr(err))
	}

	return auth.New(
		&auth.UseCases{
			Jwt:   jwt,
			Coder: coder,
		},
		&auth.Storages{
			User: storage.Users,
		},
	)
}
