package user

import (
	"time"

	"github.com/gosuit/e"
)

const (
	refreshExpires = 72 * time.Hour
	accessExpires  = 1 * time.Hour
	codeLength     = 6
)

var (
	conflictErr = e.New("User with this email already exist.", e.Conflict)
	badCodeErr  = e.New("Your activation code is wrong.", e.BadInput)
	badPassErr  = e.New("Incorrect password.", e.Forbidden)
)
