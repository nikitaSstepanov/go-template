package entity

import (
	"encoding/json"

	"github.com/nikitaSstepanov/tools/client/pg"
)

type User struct {
	Id       uint64 `redis:"id"`
	Email    string `redis:"email"`
	Name     string `redis:"name"`
	Password string `redis:"password"`
	Role     string `redis:"role"`
	Age      int    `redis:"age"`
	Verified bool   `redis:"verified"`
}

func (u *User) MarshalBinary() ([]byte, error) {
	return json.Marshal(u)
}

func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

func (u *User) Scan(r pg.Row) error {
	return r.Scan(
		&u.Id,
		&u.Email,
		&u.Name,
		&u.Password,
		&u.Role,
		&u.Age,
		&u.Verified,
	)
}
