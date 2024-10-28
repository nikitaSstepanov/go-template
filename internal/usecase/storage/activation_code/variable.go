package activation_code

import (
	"fmt"
	"time"

	e "github.com/nikitaSstepanov/tools/error"
)

const (
	redisExpires = 5 * time.Minute
)

var (
	internalErr = e.New("Something going wrong...", e.Internal)
	notFoundErr = e.New("This code wasn`t found.", e.NotFound)
)

func redisKey(userId uint64) string {
	return fmt.Sprintf("codes:%d", userId)
}
