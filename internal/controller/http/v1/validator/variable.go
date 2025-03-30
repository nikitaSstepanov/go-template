package validator

import (
	"github.com/gosuit/e"
)

const (
	Password Arg = iota
)

var (
	lenErr = e.New("Bad string length", e.BadInput)
)
