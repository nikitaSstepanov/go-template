package tests

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gavv/httpexpect/v2"
)

func TestLogoutAccount(t *testing.T) {
	url := testCfg.ToURL()
	e := httpexpect.Default(t, url)

	_, token := createUser(e)

	tests := []struct {
		TestName          string
		Token             string
		Status            int
		WithoutAuthHeader bool
		IsError           bool
	}{
		{
			TestName: "Success",
			Token:    fmt.Sprintf("Bearer %s", token),
			Status:   http.StatusOK,
		},
		{
			TestName: "Invalid token",
			Token:    fmt.Sprintf("Bearer %s", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9"),
			Status:   http.StatusUnauthorized,
			IsError:  true,
		},
		{
			TestName:          "No Authorization header",
			Status:            http.StatusUnauthorized,
			WithoutAuthHeader: true,
			IsError:           true,
		},
		// TODO: Add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.TestName, func(t *testing.T) {
			if tc.IsError {
				if tc.WithoutAuthHeader {
					e.POST("/auth/logout").Expect().Status(tc.Status).JSON().Object().ContainsKey("error")
				} else {
					e.POST("/auth/logout").WithHeader("Authorization", tc.Token).Expect().Status(tc.Status).JSON().Object().
						ContainsKey("error")
				}
			} else {
				obj := e.POST("/auth/logout").WithHeader("Authorization", tc.Token).Expect().Status(tc.Status).JSON().Object()
				obj.ContainsKey("message")
			}
		})
	}
}
