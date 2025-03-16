package account

import (
	"context"
	"time"

	"app/internal/entity"
	e "github.com/nikitaSstepanov/tools/error"
)

const (
	refreshExpires = 72 * time.Hour
	accessExpires  = 1 * time.Hour
	codeLength     = 6
	cost           = 10
)

var (
	conflictErr = e.New("User with this email already exist", e.Conflict)
	badCodeErr  = e.New("Your activation code is wrong.", e.BadInput)
	badPassErr  = e.New("Incorrect password", e.Forbidden)
	internalErr = e.New("Something going wrong...", e.Internal)
)

type UserStorage interface {
	GetById(ctx context.Context, id uint64) (*entity.User, e.Error)
	GetByEmail(ctx context.Context, email string) (*entity.User, e.Error)
	Create(ctx context.Context, user *entity.User) e.Error
	Update(ctx context.Context, user *entity.User) e.Error
	Verify(ctx context.Context, user *entity.User) e.Error
	Delete(ctx context.Context, user *entity.User) e.Error
}

type TokenStorage interface {
	Set(ctx context.Context, token *entity.Token) e.Error
}

type CodeStorage interface {
	Get(ctx context.Context, userId uint64) (*entity.ActivationCode, e.Error)
	Set(ctx context.Context, code *entity.ActivationCode) e.Error
	Del(ctx context.Context, userId uint64) e.Error
}

type JwtUseCase interface {
	GenerateToken(id uint64, role string, expires time.Duration, isRefresh bool) (string, e.Error)
}

type MailUseCase interface {
	SendActivation(to string, code string) e.Error
}
