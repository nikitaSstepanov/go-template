package tests

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect/v2"
	"app/internal/usecase/storage/activation_code"
	"app/internal/usecase/storage/user"
	"github.com/nikitaSstepanov/tools/client/pg"
	"github.com/nikitaSstepanov/tools/client/redis"
	"github.com/nikitaSstepanov/tools/sl"
)

func TestVerifyAccount(t *testing.T) {
	url := testCfg.ToURL()
	e := httpexpect.Default(t, url)

	ctx := context.TODO()

	account, token := createUser(e)

	Rclient := connectToRedis(t)
	Pclient := connectToPostgres(t)

	storage := user.New(Pclient, Rclient)
	codeStorage := activation_code.New(Rclient)

	UserData, err := storage.GetByEmail(ctx, account.Email)
	if err != nil {
		t.Fatal("Faled to get user data from storage", err.SlErr())
	}

	code, err := codeStorage.Get(ctx, UserData.Id)
	if err != nil {
		t.Fatal("Faled to get activation code from storage", err.SlErr())
	}

	tests := []struct {
		TestName      string
		Token         string
		WithoutHeader bool
		Code          string
		Status        int
		IsError       bool
	}{
		{
			TestName: "Success",
			Token:    fmt.Sprintf("Bearer %s", token),
			Code:     code.Code,
			Status:   http.StatusOK,
		},
		{
			TestName: "Invalid code",
			Token:    fmt.Sprintf("Bearer %s", token),
			Code:     strconv.Itoa(int(gofakeit.Uint16())),
			Status:   http.StatusBadRequest,
			IsError:  true,
		},
		{
			TestName:      "Without authorization header",
			Code:          code.Code,
			Status:        http.StatusUnauthorized,
			IsError:       true,
			WithoutHeader: true,
		},
		{
			TestName: "Invalid token",
			Token:    fmt.Sprintf("Bearer %s", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"),
			Code:     code.Code,
			Status:   http.StatusUnauthorized,
			IsError:  true,
		},

		//TODO: Add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.TestName, func(t *testing.T) {
			if tc.IsError {
				if tc.WithoutHeader {
					e.GET(fmt.Sprintf("/verify/confirm/%s", tc.Code)).
						Expect().Status(tc.Status).JSON().
						Object().ContainsKey("error")
				} else {
					e.GET(fmt.Sprintf("/verify/confirm/%s", tc.Code)).
						WithHeader("Authorization", tc.Token).
						Expect().Status(tc.Status).JSON().
						Object().ContainsKey("error")
				}
			} else {
				e.GET(fmt.Sprintf("/verify/confirm/%s", tc.Code)).
					WithHeader("Authorization", tc.Token).Expect().Status(tc.Status).
					JSON().
					Object().ContainsKey("message")
			}
		})
	}
}

func connectToPostgres(t *testing.T) pg.Client {
	client, err := pg.ConnectToDb(context.TODO(), &testCfg.Postgres)
	if err != nil {
		t.Fatal("Failed to connect to postgres", sl.ErrAttr(err))
	}

	return client
}

func connectToRedis(t *testing.T) redis.Client {
	client, err := redis.ConnectToRedis(context.TODO(), &testCfg.Redis)
	if err != nil {
		t.Fatal("Failed to connect to redis", sl.ErrAttr(err))
	}

	return client
}
