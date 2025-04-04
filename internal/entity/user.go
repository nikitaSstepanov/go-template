package entity

import (
	"app/internal/entity/types"
	"encoding/json"

	"github.com/gosuit/pg"
)

type User struct {
	Id       uint64     `redis:"id"`
	Email    string     `redis:"email"`
	Name     string     `redis:"name"`
	Password string     `redis:"password"`
	Age      int        `redis:"age"`
	Role     types.Role `redis:"role"`
	Verified bool       `redis:"verified"`
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
		&u.Age,
		&u.Role,
		&u.Verified,
	)
}
