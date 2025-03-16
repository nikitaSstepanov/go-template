package account

import (
	"context"
	"net/http"

	resp "app/internal/controller/response"
	"app/internal/entity"

	"github.com/gosuit/e"
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

	verifiedMsg = resp.NewMessage("Verified.")
	updatedMsg  = resp.NewMessage("Updated.")
	okMsg       = resp.NewMessage("Ok.")
)

type AccountUseCase interface {
	Get(ctx context.Context, userId uint64) (*entity.User, e.Error)
	Create(ctx context.Context, user *entity.User) (*entity.Tokens, e.Error)
	Update(ctx context.Context, user *entity.User, pass string) e.Error
	Verify(ctx context.Context, id uint64, code string) e.Error
	ResendCode(ctx context.Context, userId uint64) e.Error
	Delete(ctx context.Context, user *entity.User) e.Error
}
