package mail

import "github.com/gosuit/e"

const (
	activateSubject = "Email verification"

	htmlType = "text/html"
)

var (
	internalErr = e.New("Something going wrong...", e.Internal)
)
