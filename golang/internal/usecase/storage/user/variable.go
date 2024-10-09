package user

import (
	"time"

	e "github.com/nikitaSstepanov/tools/error"
)

const (
	redisExpires = 3 * time.Hour
	usersTable   = "users"
)

var (
	internalErr = e.New("Something going wrong...", e.Internal)
	notFoundErr = e.New("This user wasn`t found.", e.NotFound)
)
