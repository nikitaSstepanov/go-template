package token

import (
	"fmt"
	"time"

	"github.com/gosuit/e"
)

const (
	redisExpires = 72 * time.Hour
)

var (
	internalErr = e.New("Something going wrong...", e.Internal)
	notFoundErr = e.New("Your token wasn`t found.", e.NotFound)
)

func redisKey(userId uint64) string {
	return fmt.Sprintf("tokens:%d", userId)
}
