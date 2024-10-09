package token

import (
	"fmt"
	"time"

	e "github.com/nikitaSstepanov/tools/error"
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
