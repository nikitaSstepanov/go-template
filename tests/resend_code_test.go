package tests

import (
	"fmt"
	"net/http"
	"net/url"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestResendCode(t *testing.T) {
	u := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   "/api/v1/account",
	}
	e := httpexpect.Default(t, u.String())

	_, token := createUser(e)

	tests := []struct {
		TestName      string
		Token         string
		WithoutHeader bool
		Status        int
		IsError       bool
	}{
		{
			TestName: "Success",
			Token:    fmt.Sprintf("Bearer %s", token),
			Status:   http.StatusOK,
		},
		{
			TestName:      "Without authorization header",
			Status:        http.StatusUnauthorized,
			IsError:       true,
			WithoutHeader: true,
		},
		{
			TestName: "Invalid token",
			Token:    fmt.Sprintf("Bearer %s", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"),
			Status:   http.StatusUnauthorized,
			IsError:  true,
		},

		//TODO: add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.TestName, func(t *testing.T) {
			if tc.IsError {
				if tc.WithoutHeader {
					e.GET("/verify/resend").
						Expect().Status(tc.Status).JSON().
						Object().ContainsKey("error")
				} else {
					e.GET("/verify/resend").
						WithHeader("Authorization", tc.Token).
						Expect().Status(tc.Status).JSON().
						Object().ContainsKey("error")
				}
			} else {
				e.GET("/verify/resend").
					WithHeader("Authorization", tc.Token).Expect().Status(tc.Status).
					JSON().
					Object().ContainsKey("message")
			}
		})
	}
}
