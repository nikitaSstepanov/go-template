package user

import (
	"app/internal/entity"

	"github.com/gosuit/e"
	"github.com/gosuit/lec"
	"github.com/gosuit/utils/coder"
	"github.com/gosuit/utils/generator"
)

type User struct {
	user  UserStorage
	code  CodeStorage
	jwt   JwtUseCase
	mail  MailUseCase
	coder coder.Coder
}

func New(uc *UseCases, store *Storages) *User {
	return &User{
		user:  store.User,
		code:  store.Code,
		jwt:   uc.Jwt,
		mail:  uc.Mail,
		coder: uc.Coder,
	}
}

func (u *User) Get(ctx lec.Context, userId uint64) (*entity.User, e.Error) {
	return u.user.GetById(ctx, userId)
}

func (u *User) Create(ctx lec.Context, user *entity.User) (*entity.Tokens, e.Error) {
	candidate, err := u.user.GetByEmail(ctx, user.Email)
	if err != nil && err.GetCode() != e.NotFound {
		return nil, err
	}

	if candidate != nil {
		return nil, conflictErr.WithCtx(ctx)
	}

	hash, hashErr := u.coder.Hash(user.Password)
	if hashErr != nil {
		return nil, e.InternalErr.
			WithCtx(ctx).WithErr(hashErr)
	}

	user.Password = hash

	err = u.user.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	go u.sendCode(ctx, user)

	access, err := u.jwt.GenerateToken(user, accessExpires, false)
	if err != nil {
		return nil, err
	}

	refresh, err := u.jwt.GenerateToken(user, refreshExpires, true)
	if err != nil {
		return nil, err
	}

	tokens := &entity.Tokens{
		Access:  access,
		Refresh: refresh,
	}

	return tokens, nil
}

func (u *User) Update(ctx lec.Context, user *entity.User, pass string) e.Error {
	if user.Email != "" {
		candidate, err := u.user.GetByEmail(ctx, user.Email)
		if err != nil && err.GetCode() != e.NotFound {
			return err
		}

		if candidate != nil {
			return conflictErr.WithCtx(ctx)
		}

		user.Verified = false

		err = u.user.Verify(ctx, user)
		if err != nil {
			return err
		}

		go u.sendCode(ctx, user)
	}

	if user.Password != "" {
		old, err := u.user.GetById(ctx, user.Id)
		if err != nil {
			return err
		}

		if err := u.coder.CompareHash(old.Password, pass); err != nil {
			return badPassErr.WithErr(err)
		}

		hash, hashErr := u.coder.Hash(user.Password)
		if hashErr != nil {
			return e.InternalErr.
				WithCtx(ctx).WithErr(hashErr)
		}

		user.Password = hash
	}

	return u.user.Update(ctx, user)
}

func (u *User) Verify(ctx lec.Context, id uint64, code string) e.Error {
	acode, err := u.code.Get(ctx, id)
	if err != nil {
		return err
	}

	if acode.Code != code {
		return badCodeErr.WithCtx(ctx)
	}

	err = u.code.Del(ctx, id)
	if err != nil {
		return err
	}

	user := &entity.User{
		Id:       id,
		Verified: true,
	}

	return u.user.Verify(ctx, user)
}

func (u *User) ResendCode(ctx lec.Context, userId uint64) e.Error {
	user, err := u.user.GetById(ctx, userId)
	if err != nil {
		return err
	}

	go u.sendCode(ctx, user)

	return nil
}

func (u *User) Delete(ctx lec.Context, user *entity.User) e.Error {
	_, err := u.user.GetById(ctx, user.Id)
	if err != nil {
		return err
	}

	return u.user.Delete(ctx, user)
}

func (u *User) sendCode(ctx lec.Context, user *entity.User) {
	log := ctx.Logger()

	_, err := u.code.Get(ctx, user.Id)
	if err != nil && err.GetCode() != e.NotFound {
		log.Error("Failed to get code from CodeStorage", err.SlErr())
	}

	if err == nil {
		err = u.code.Del(ctx, user.Id)
		if err != nil {
			log.Error("Failed to delete code from CodeStorage", err.SlErr())
		}
	}

	code := generator.GetRandomNum(codeLength)

	acode := &entity.ActivationCode{
		Code:   code,
		UserId: user.Id,
	}

	err = u.code.Set(ctx, acode)
	if err != nil {
		log.Error("Failed to set code in CodeStorage", err.SlErr())
	}

	err = u.mail.SendActivation(user.Email, code)
	if err != nil {
		log.Error("Failed to send code with smtp", err.SlErr())
	}
}
