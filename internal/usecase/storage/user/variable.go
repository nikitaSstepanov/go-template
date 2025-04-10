package user

import (
	"time"

	"github.com/gosuit/e"
)

const (
	redisExpires = 3 * time.Hour
	usersTable   = "users"
)

var (
	notFoundErr = e.New("This user wasn`t found.", e.NotFound)
)
