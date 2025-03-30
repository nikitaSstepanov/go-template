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
		// TODO: Add more test cases
	}

	for _, tc := range tests {
		t.Run(tc.TestName, func(t *testing.T) {
			e.POST("/auth/logout").Expect().Status(tc.Status)
		})
	}
}
