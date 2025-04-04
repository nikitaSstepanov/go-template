package auth

import (
	"time"

	"github.com/gosuit/e"
)

const (
	refreshExpires = 72 * time.Hour
	accessExpires  = 1 * time.Hour
)

var (
	badDataErr = e.New("Incorrect email or password.", e.Unauthorize)
	
)
