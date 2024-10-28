package entity

import "encoding/json"

type ActivationCode struct {
	Code   string `redis:"code"`
	UserId uint64 `redis:"userId"`
}

func (c *ActivationCode) MarshalBinary() ([]byte, error) {
	return json.Marshal(c)
}

func (c *ActivationCode) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}
