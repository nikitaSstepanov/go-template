package entity

import "encoding/json"

type Tokens struct {
	Access  string
	Refresh string
}

type Token struct {
	Token  string `redis:"token"`
	UserId uint64 `redis:"userId"`
}

func (t *Token) MarshalBinary() ([]byte, error) {
	return json.Marshal(t)
}

func (t *Token) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, t)
}
