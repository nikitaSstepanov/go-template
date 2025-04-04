package account

import (
	resp "app/internal/controller/response"

	"github.com/gosuit/e"
)

var (
	badReqErr = e.New("Incorrect data.", e.BadInput)

	verifiedMsg = resp.NewMessage("Verified.")
	updatedMsg  = resp.NewMessage("Updated.")
	okMsg       = resp.NewMessage("Ok.")
)
