package tests

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"net/url"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect/v2"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/dto"
)

func TestGetAccount(t *testing.T) {

	u := url.URL{
		Scheme: "http",
		Host:   host,
		Path:   "/api/v1/account",
	}

	e := httpexpect.Default(t, u.String())

	user, token := createUser(e)

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
			TestName:      "No Authorization header",
			WithoutHeader: true,
			Status:        http.StatusUnauthorized,
			IsError:       true,
		},
		{
			TestName: "Token is not bearer",
			Token:    token,
			Status:   http.StatusUnauthorized,
			IsError:  true,
		},
		{
			TestName: "Invalig token",
			Token:    fmt.Sprintf("Bearer %s", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"),
			Status:   http.StatusUnauthorized,
			IsError:  true,
		},

		//TODO: Add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.TestName, func(t *testing.T) {
			if tc.IsError {
				if tc.WithoutHeader {
					e.GET("/").Expect().Status(tc.Status).JSON().Object().ContainsKey("error")
				} else {
					e.GET("/").WithHeader("Authorization", tc.Token).Expect().Status(tc.Status).JSON().Object().
						ContainsKey("error")
				}
			} else {
				e.GET("/").WithHeader("Authorization", tc.Token).Expect().Status(tc.Status).JSON().Object().
					HasValue("email", user.Email).
					HasValue("name", user.Name).
					HasValue("age", user.Age).
					ContainsKey("id")
			}
		})
	}
}

func createUser(e *httpexpect.Expect) (dto.CreateUser, string) {
	user := dto.CreateUser{
		Email:    genRandEmail(),
		Name:     gofakeit.Name(),
		Age:      rand.IntN(100),
		Password: gofakeit.Password(true, true, true, true, false, 10),
	}

	val := e.POST("/new").WithJSON(user).Expect().Status(http.StatusOK).JSON().Object().Iter()["token"]
	token := val.String().Raw()
	return user, token
}
