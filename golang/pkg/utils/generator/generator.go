package generator

import "math/rand"

const (
	lenOfSymbols = 10
)

var (
	symbols = []rune("1234567890")
)

func GetRandomNum(length int) string {
	res := make([]rune, length)

	for i := range res {
		res[i] = symbols[rand.Intn(lenOfSymbols)]
	}

	return string(res)
}
