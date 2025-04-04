package activation_code

import "fmt"

func redisKey(userId uint64) string {
	return fmt.Sprintf("activation_codes:%d", userId)
}
