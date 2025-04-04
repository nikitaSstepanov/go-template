package dto

type Account struct {
	Id       uint64 `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Verified bool   `json:"verified"`
}

type CreateUser struct {
	Email    string `json:"email"    validate:"required,email"`
	Name     string `json:"name"     validate:"required,min=1"`
	Password string `json:"password" validate:"required,min=8,max=50,password"`
	Age      int    `json:"age"      validate:"gte=0,lte=200"`
}

type UpdateUser struct {
	Email       string `json:"email"        validate:"required,email"`
	Name        string `json:"name"         validate:"required,min=1"`
	Password    string `json:"password"     validate:"omitempty,min=8,max=50,password"`
	OldPassword string `json:"old_password" validate:"omitempty,min=8,max=50,password"`
	Age         int    `json:"age"          validate:"required,gte=0,lte=200"`
}

type SetRole struct {
	Id   uint64 `json:"user_id" validate:"required"`
	Role string `json:"role"    validate:"required,oneof=USER ADMIN"`
}
