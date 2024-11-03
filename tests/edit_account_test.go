package tests

import (
	"fmt"
	"math/rand/v2"
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/gavv/httpexpect/v2"
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/dto"
)

func TestEditAccount(t *testing.T) {
	url := testCfg.ToURL()
	e := httpexpect.Default(t, url)

	user, token := createUser(e)

	tests := []struct {
		TestName    string
		Email       string
		Name        string
		Age         int
		OldPassword string
		Password    string
		Token       string
		Status      int
		IsError     bool
	}{
		{
			TestName: "Invalid Password",
			Email:    genRandEmail(),
			Name:     gofakeit.Name(),
			Age:      rand.IntN(100),
			Token:    fmt.Sprintf("Bearer %s", token),
			Password: gofakeit.Password(true, true, true, true, false, 5),
			Status:   http.StatusBadRequest,
			IsError:  true,
		},
		{
			TestName: "Invalid Age",
			Email:    genRandEmail(),
			Name:     gofakeit.Name(),
			Age:      -50,
			Token:    fmt.Sprintf("Bearer %s", token),
			Password: gofakeit.Password(true, true, true, true, false, 10),
			Status:   http.StatusBadRequest,
			IsError:  true,
		},
		{
			TestName: "Invalid Email",
			Email:    "adhfianfdfgdg",
			Name:     gofakeit.Name(),
			Age:      rand.IntN(100),
			Token:    fmt.Sprintf("Bearer %s", token),
			Password: gofakeit.Password(true, true, true, true, false, 10),
			Status:   http.StatusBadRequest,
			IsError:  true,
		},
		{
			TestName:    "Unauthorized",
			Email:       genRandEmail(),
			Name:        gofakeit.Name(),
			Age:         rand.IntN(100),
			OldPassword: user.Password,
			Password:    gofakeit.Password(true, true, true, true, false, 10),
			Status:      http.StatusUnauthorized,
			IsError:     true,
		},
		{
			TestName:    "Wrong Old Passwaord",
			Email:       genRandEmail(),
			Name:        gofakeit.Name(),
			Age:         rand.IntN(100),
			OldPassword: gofakeit.Password(true, true, true, true, false, 10),
			Password:    gofakeit.Password(true, true, true, true, false, 10),
			Token:       fmt.Sprintf("Bearer %s", token),
			Status:      http.StatusForbidden,
			IsError:     true,
		},
		{
			TestName:    "Success",
			Email:       genRandEmail(),
			Name:        gofakeit.Name(),
			Age:         rand.IntN(100),
			Token:       fmt.Sprintf("Bearer %s", token),
			OldPassword: user.Password,
			Password:    gofakeit.Password(true, true, true, true, false, 10),
			Status:      http.StatusOK,
		},
		{
			TestName: "Empty fields",
			Email:    genRandEmail(),
			Token:    fmt.Sprintf("Bearer %s", token),
			Status:   http.StatusOK,
		},
		//TODO: Add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.TestName, func(t *testing.T) {
			obj := e.PATCH("/edit").WithHeader("Authorization", tc.Token).WithJSON(dto.UpdateUser{
				Email:       tc.Email,
				Name:        tc.Name,
				Age:         tc.Age,
				OldPassword: tc.OldPassword,
				Password:    tc.Password,
			}).Expect().Status(tc.Status).JSON().Object()
			if tc.IsError {
				obj.ContainsKey("error")
			} else {
				obj.ContainsKey("message")
			}
		})
	}
}
