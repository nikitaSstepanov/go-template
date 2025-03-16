package mail

import (
	"github.com/gosuit/e"
	gomail "github.com/gosuit/mail"
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
		return e.InternalErr.WithErr(err)
	}

	return nil
}
