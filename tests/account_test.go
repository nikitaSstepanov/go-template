package tests

import (
	"math/rand/v2"
	"net/http"
	"net/url"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect/v2"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/dto"
)

const (
	host = "localhost:8080"
)

func TestNewUser(t *testing.T) {
	u := url.URL{
		Scheme: "http",
		Host:   host,
	}

	e := httpexpect.Default(t, u.String())

	e.POST("/api/v1/account/new").WithJSON(dto.CreateUser{
		Email:    gofakeit.Email(),
		Name:     gofakeit.Name(),
		Age:      rand.IntN(100),
		Password: gofakeit.Password(true, true, true, true, false, 10),
	}).
		Expect().Status(http.StatusOK).JSON().Object().ContainsKey("token")

}
