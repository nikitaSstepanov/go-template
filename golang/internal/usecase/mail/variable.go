package mail

import e "github.com/nikitaSstepanov/tools/error"

const (
	activateSubject = "Email verification"

	htmlType = "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";"
)

var (
	internalErr = e.New("Something going wrong...", e.Internal)
)
