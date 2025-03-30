package resp

type Message struct {
	Message string `json:"message"`
}

func NewMessage(msg string) *Message {
	return &Message{Message: msg}
}

// JsonError use only for doc and represent e.JsonError
type JsonError struct {
	Error string `json:"error"`
}
