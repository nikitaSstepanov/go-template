package account

import (
	"context"
	"net/http"

	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/dto"
	"github.com/nikitaSstepanov/templates/golang/internal/entity"
	e "github.com/nikitaSstepanov/tools/error"
)

const (
	ok      = http.StatusOK
	created = http.StatusCreated
	badReq  = http.StatusBadRequest
	unauth  = http.StatusUnauthorized

	cookieName = "refreshToken"
	cookieAge  = 259200
	cookiePath = "/"
	cookieHost = "localhost"
)

var (
	badReqErr = e.New("Incorrect data.", e.BadInput)

	updatedMsg  = dto.NewMessage("Updated.")
	verifiedMsg = dto.NewMessage("Verified.")
	okMsg       = dto.NewMessage("Ok.")
)

type AccountUseCase interface {
	Get(ctx context.Context, userId uint64) (*entity.User, e.Error)
	Create(ctx context.Context, user *entity.User) (*entity.Tokens, e.Error)
	Update(ctx context.Context, user *entity.User, pass string) e.Error
	Verify(ctx context.Context, id uint64, code string) e.Error
	ResendCode(ctx context.Context, userId uint64) e.Error
	Delete(ctx context.Context, user *entity.User) e.Error
}
