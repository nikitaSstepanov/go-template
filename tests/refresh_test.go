package tests

import (
	"math/rand/v2"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect/v2"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/dto"
)

func TestRefresh(t *testing.T) {
	url := testCfg.ToURL()

	e := httpexpect.Default(t, url)

	t.Run("Without cookie", func(t *testing.T) {
		e.GET("auth/refresh").Expect().Status(http.StatusBadRequest).
			JSON().Object().ContainsKey("error")
	})

	user := dto.CreateUser{
		Email:    genRandEmail(),
		Name:     gofakeit.Name(),
		Age:      rand.IntN(100),
		Password: gofakeit.Password(true, true, true, true, false, 10),
	}

	val := e.POST("/new").WithJSON(user).Expect().Status(http.StatusOK).Cookie("refreshToken").Value()
	token := val.Raw()

	test := []struct {
		TestName      string
		RefreshToken  string
		Status        int
		WithoutCookie bool
		IsError       bool
	}{
		{
			TestName:     "Success",
			Status:       http.StatusOK,
			RefreshToken: token,
		},
		{
			TestName:     "Invalid token",
			Status:       http.StatusUnauthorized,
			RefreshToken: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
			IsError:      true,
		},
		{
			TestName: "Empty refresh token",
			Status:   http.StatusUnauthorized,
			IsError:  true,
		},
	}

	for _, tc := range test {
		t.Run(tc.TestName, func(t *testing.T) {
			if tc.IsError {
				e.GET("auth/refresh").WithCookie("refreshToken", tc.RefreshToken).Expect().Status(tc.Status).
					JSON().Object().ContainsKey("error")
			} else {
				e.GET("auth/refresh").WithCookie("refreshToken", tc.RefreshToken).Expect().Status(tc.Status).
					JSON().Object().ContainsKey("token")
			}
		})
	}

}
