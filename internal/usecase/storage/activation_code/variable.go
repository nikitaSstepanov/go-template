package activation_code

import (
	"time"

	"github.com/gosuit/e"
)

const (
	redisExpires = 5 * time.Minute
)

var (
	notFoundErr = e.New("This code wasn`t found.", e.NotFound)
)
