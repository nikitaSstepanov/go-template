package tests

import (
	"math/rand/v2"
	"net/http"
	"strings"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect/v2"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/dto"
)

func genRandEmail() string {
	domens := []string{"@mail.com", "@yandex.com", "@gmail.com",
		"@hotmal.com", "@email.com"}
	return strings.Split(gofakeit.Email(), "@")[0] + gofakeit.RandString(domens)
}

func TestCreate(t *testing.T) {
	url := testCfg.ToURL()
	e := httpexpect.Default(t, url)

	tests := []struct {
		TestName string
		Email    string
		Name     string
		Age      int
		Password string
		Status   int
		IsError  bool
	}{
		{
			TestName: "Success",
			Email:    genRandEmail(),
			Name:     gofakeit.Name(),
			Age:      rand.IntN(100),
			Password: gofakeit.Password(true, true, true, true, false, 10),
			Status:   http.StatusOK,
		},
		{
			TestName: "Invalid Password",
			Email:    genRandEmail(),
			Name:     gofakeit.Name(),
			Age:      rand.IntN(100),
			Password: gofakeit.Password(true, true, true, true, false, 5),
			Status:   http.StatusBadRequest,
			IsError:  true,
		},
		{
			TestName: "Invalid Age",
			Email:    genRandEmail(),
			Name:     gofakeit.Name(),
			Age:      -50,
			Password: gofakeit.Password(true, true, true, true, false, 10),
			Status:   http.StatusBadRequest,
			IsError:  true,
		},
		{
			TestName: "Invalid Email",
			Email:    "adhfiaasaassdanfgdg",
			Name:     gofakeit.Name(),
			Age:      rand.IntN(100),
			Password: gofakeit.Password(true, true, true, true, false, 10),
			Status:   http.StatusBadRequest,
			IsError:  true,
		},

		//TODO: Add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.TestName, func(t *testing.T) {
			obj := e.POST("/new").WithJSON(dto.CreateUser{
				Email:    tc.Email,
				Name:     tc.Name,
				Age:      tc.Age,
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
