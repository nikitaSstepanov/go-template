package auth

import (
	"net/http"

	"github.com/gosuit/e"
)

const (
	ok      = http.StatusOK
	created = http.StatusCreated
	badReq  = http.StatusBadRequest
	unauth  = http.StatusUnauthorized
)

var (
	badReqErr = e.New("Incorrect data.", e.BadInput)
)
