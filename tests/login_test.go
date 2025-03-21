package tests

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect/v2"
	"app/internal/controller/http/v1/dto"
)

func TestLogin(t *testing.T) {
	url := testCfg.ToURL()
	e := httpexpect.Default(t, url)

	user, _ := createUser(e)

	tests := []struct {
		TestName string
		Email    string
		Password string
		Status   int
		IsError  bool
	}{
		{
			TestName: "Success",
			Email:    user.Email,
			Password: user.Password,
			Status:   http.StatusOK,
		},
		{
			TestName: "Incorrect password",
			Email:    user.Email,
			Password: gofakeit.Password(true, true, true, true, false, 10),
			Status:   http.StatusUnauthorized,
			IsError:  true,
		},
		{
			TestName: "User not found",
			Email:    genRandEmail(),
			Password: gofakeit.Password(true, true, true, true, false, 10),
			Status:   http.StatusNotFound,
			IsError:  true,
		},
		// TODO: Add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.TestName, func(t *testing.T) {
			obj := e.POST("/auth/login").WithJSON(dto.Login{
				Email:    tc.Email,
				Password: tc.Password,
			}).Expect().Status(tc.Status).JSON().Object()

			if tc.IsError {
				obj.ContainsKey("error")
			} else {
				obj.ContainsKey("token")
			}
		})
	}
}
