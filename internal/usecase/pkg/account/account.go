package account

import (
	"app/internal/entity"
	"context"

	"github.com/gosuit/e"
	"github.com/gosuit/sl"
	"github.com/gosuit/utils/generator"
	"golang.org/x/crypto/bcrypt"
)

type Account struct {
	user  UserStorage
	token TokenStorage
	code  CodeStorage
	jwt   JwtUseCase
	mail  MailUseCase
}

func New(user UserStorage, token TokenStorage, code CodeStorage, jwt JwtUseCase, mail MailUseCase) *Account {
	return &Account{
		user:  user,
		token: token,
		code:  code,
		jwt:   jwt,
		mail:  mail,
	}
}

func (a *Account) Get(ctx context.Context, userId uint64) (*entity.User, e.Error) {
	return a.user.GetById(ctx, userId)
}

func (a *Account) Create(ctx context.Context, user *entity.User) (*entity.Tokens, e.Error) {
	candidate, err := a.user.GetByEmail(ctx, user.Email)
	if err != nil && err.GetCode() != e.NotFound {
		return nil, err
	}

	if candidate != nil {
		return nil, conflictErr
	}

	hash, hashErr := hashPassword(user.Password)
	if hashErr != nil {
		return nil, internalErr.WithErr(err)
	}

	user.Password = hash

	err = a.user.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	go a.sendCode(ctx, user)

	var tokens entity.Tokens

	access, err := a.jwt.GenerateToken(user.Id, "USER", accessExpires, false)
	if err != nil {
		return nil, err
	}

	refresh, err := a.jwt.GenerateToken(user.Id, "USER", refreshExpires, true)
	if err != nil {
		return nil, err
	}

	token := &entity.Token{
		Token:  refresh,
		UserId: user.Id,
	}

	err = a.token.Set(ctx, token)
	if err != nil {
		return nil, err
	}

	tokens.Access = access
	tokens.Refresh = refresh

	return &tokens, nil
}

func (a *Account) Update(ctx context.Context, user *entity.User, pass string) e.Error {
	if user.Email != "" {
		candidate, err := a.user.GetByEmail(ctx, user.Email)
		if err != nil && err.GetCode() != e.NotFound {
			return err
		}

		if candidate != nil {
			return conflictErr
		}

		user.Verified = false

		err = a.user.Verify(ctx, user)
		if err != nil {
			return err
		}

		go a.sendCode(ctx, user)

	}

	if user.Password != "" {
		old, err := a.user.GetById(ctx, user.Id)
		if err != nil {
			return err
		}

		if err := checkPassword(old.Password, pass); err != nil {
			return badPassErr.WithErr(err)
		}

		hash, hashErr := hashPassword(user.Password)
		if hashErr != nil {
			return internalErr.WithErr(err)
		}

		user.Password = hash
	}

	return a.user.Update(ctx, user)
}

func (a *Account) Verify(ctx context.Context, id uint64, code string) e.Error {
	acode, err := a.code.Get(ctx, id)
	if err != nil {
		return err
	}

	if acode.Code != code {
		return badCodeErr
	}

	err = a.code.Del(ctx, id)
	if err != nil {
		return err
	}

	user := &entity.User{
		Id:       id,
		Verified: true,
	}

	return a.user.Verify(ctx, user)
}

func (a *Account) ResendCode(ctx context.Context, userId uint64) e.Error {
	user, err := a.user.GetById(ctx, userId)
	if err != nil {
		return err
	}

	go a.sendCode(ctx, user)

	return nil
}

func (a *Account) Delete(ctx context.Context, user *entity.User) e.Error {
	toDel, err := a.user.GetById(ctx, user.Id)
	if err != nil {
		return err
	}

	if err := checkPassword(toDel.Password, user.Password); err != nil {
		return badPassErr.WithErr(err)
	}

	return a.user.Delete(ctx, user)
}

func (a *Account) sendCode(ctx context.Context, user *entity.User) {
	log := sl.L(ctx)
	_, err := a.code.Get(ctx, user.Id)
	if err != nil && err.GetCode() != e.NotFound {
		log.Error("Failed to get code from CodeStorage", err.SlErr())
	}

	if err == nil {
		err = a.code.Del(ctx, user.Id)
		if err != nil {
			log.Error("Failed to delete code from CodeStorage", err.SlErr())

		}
	}

	code := generator.GetRandomNum(codeLength)

	acode := &entity.ActivationCode{
		Code:   code,
		UserId: user.Id,
	}

	err = a.code.Set(ctx, acode)
	if err != nil {
		log.Error("Failed to set code in CodeStorage", err.SlErr())
	}

	err = a.mail.SendActivation(user.Email, code)
	if err != nil {
		log.Error("Failed to send code with smtp", err.SlErr())
	}
}

func checkPassword(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func hashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}
