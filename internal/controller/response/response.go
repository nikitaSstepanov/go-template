package resp

type Message struct {
	Message string `json:"message"`
}

func NewMessage(msg string) *Message {
	return &Message{Message: msg}
}
