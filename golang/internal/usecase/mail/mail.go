package mail

import (
	"fmt"

	gomail "github.com/nikitaSstepanov/tools/client/mail"
	e "github.com/nikitaSstepanov/tools/error"
)

type Mail struct {
	client *gomail.Client
}

func New(cfg *gomail.Config) *Mail {
	return &Mail{
		client: gomail.New(cfg),
	}
}

func (m *Mail) SendActivation(to string, code string) e.Error {
	message := activationMessage(code)

	if err := m.client.Send(to, message, activateSubject, htmlType); err != nil {
		return internalErr.WithErr(fmt.Errorf(
			"Failed to send mail, %s", err.Error(),
			))
	}

	return nil
}
