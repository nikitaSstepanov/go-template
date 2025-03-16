package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect/v2"
	"app/internal/controller/http/v1/dto"
)

func TestDeleteAccount(t *testing.T) {
	url := testCfg.ToURL()
	e := httpexpect.Default(t, url)

	user, token := createUser(e)

	tests := []struct {
		TestName      string
		Token         string
		Password      string
		WithoutHeader bool
		Status        int
		IsError       bool
	}{
		{
			TestName:      "No Authorization header",
			WithoutHeader: true,
			Password:      user.Password,
			Status:        http.StatusUnauthorized,
			IsError:       true,
		},
		{
			TestName: "Token is not bearer",
			Token:    token,
			Password: user.Password,
			Status:   http.StatusUnauthorized,
			IsError:  true,
		},
		{
			TestName: "Invalid token",
			Token:    fmt.Sprintf("Bearer %s", "eyJhbGdeiOiJRUzI1NiIsInR5cCI6IkpXVCJ9"),
			Password: user.Password,
			Status:   http.StatusUnauthorized,
			IsError:  true,
		},
		{
			TestName: "Invalid password",
			Token:    fmt.Sprintf("Bearer %s", token),
			Password: gofakeit.Password(true, true, true, true, false, 10),
			Status:   http.StatusForbidden,
			IsError:  true,
		},
		{
			TestName: "Success",
			Token:    fmt.Sprintf("Bearer %s", token),
			Password: user.Password,
			Status:   http.StatusOK,
		},

		//TODO: Add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.TestName, func(t *testing.T) {
			if tc.IsError {
				if tc.WithoutHeader {
					e.DELETE("/delete").Expect().Status(tc.Status).JSON().Object().ContainsKey("error")
				} else {
					e.DELETE("/delete").WithHeader("Authorization", tc.Token).
						WithJSON(dto.DeleteUser{Password: tc.Password}).
						Expect().Status(tc.Status).JSON().Object().
						ContainsKey("error")
				}
			} else {
				e.DELETE("/delete").WithHeader("Authorization", tc.Token).
					WithJSON(dto.DeleteUser{Password: tc.Password}).
					Expect().Status(tc.Status).JSON().Object().
					ContainsKey("message")
			}
		})
	}
}
