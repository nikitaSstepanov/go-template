package auth

import (
	"github.com/gosuit/e"
)

var (
	badReqErr = e.New("Incorrect data.", e.BadInput)
)
