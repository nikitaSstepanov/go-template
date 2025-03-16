package account

import (
	"app/internal/entity"

	"github.com/gosuit/e"
	"github.com/gosuit/lec"
	"github.com/gosuit/utils/coder"
	"github.com/gosuit/utils/generator"
)

type Account struct {
	user  UserStorage
	code  CodeStorage
	jwt   JwtUseCase
	mail  MailUseCase
	coder *coder.Coder
}

func New(uc *UseCases, store *Storages) *Account {
	return &Account{
		user:  store.User,
		code:  store.Code,
		jwt:   uc.Jwt,
		mail:  uc.Mail,
		coder: uc.Coder,
	}
}

func (a *Account) Get(ctx lec.Context, userId uint64) (*entity.User, e.Error) {
	return a.user.GetById(ctx, userId)
}

func (a *Account) Create(ctx lec.Context, user *entity.User) (*entity.Tokens, e.Error) {
	candidate, err := a.user.GetByEmail(ctx, user.Email)
	if err != nil && err.GetCode() != e.NotFound {
		return nil, err
	}

	if candidate != nil {
		return nil, conflictErr.WithCtx(ctx)
	}

	hash, hashErr := a.coder.Hash(user.Password)
	if hashErr != nil {
		return nil, e.InternalErr.
			WithCtx(ctx).WithErr(hashErr)
	}

	user.Password = hash

	err = a.user.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	go a.sendCode(ctx, user)

	access, err := a.jwt.GenerateToken(user, accessExpires, false)
	if err != nil {
		return nil, err
	}

	refresh, err := a.jwt.GenerateToken(user, refreshExpires, true)
	if err != nil {
		return nil, err
	}

	tokens := &entity.Tokens{
		Access:  access,
		Refresh: refresh,
	}

	return tokens, nil
}

func (a *Account) Update(ctx lec.Context, user *entity.User, pass string) e.Error {
	if user.Email != "" {
		candidate, err := a.user.GetByEmail(ctx, user.Email)
		if err != nil && err.GetCode() != e.NotFound {
			return err
		}

		if candidate != nil {
			return conflictErr.WithCtx(ctx)
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

		if err := a.coder.CompareHash(old.Password, pass); err != nil {
			return badPassErr.WithErr(err)
		}

		hash, hashErr := a.coder.Hash(user.Password)
		if hashErr != nil {
			return e.InternalErr.
				WithCtx(ctx).WithErr(hashErr)
		}

		user.Password = hash
	}

	return a.user.Update(ctx, user)
}

func (a *Account) Verify(ctx lec.Context, id uint64, code string) e.Error {
	acode, err := a.code.Get(ctx, id)
	if err != nil {
		return err
	}

	if acode.Code != code {
		return badCodeErr.WithCtx(ctx)
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

func (a *Account) ResendCode(ctx lec.Context, userId uint64) e.Error {
	user, err := a.user.GetById(ctx, userId)
	if err != nil {
		return err
	}

	go a.sendCode(ctx, user)

	return nil
}

func (a *Account) Delete(ctx lec.Context, user *entity.User) e.Error {
	_, err := a.user.GetById(ctx, user.Id)
	if err != nil {
		return err
	}

	return a.user.Delete(ctx, user)
}

func (a *Account) sendCode(ctx lec.Context, user *entity.User) {
	log := ctx.Logger()

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
