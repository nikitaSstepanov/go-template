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
	Email       string `json:"email"       validate:"omitempty,email"`
	Name        string `json:"name"        validate:"omitempty,min=1"`
	Password    string `json:"password"    validate:"omitempty,min=8,max=50,password"`
	OldPassword string `json:"oldPassword" validate:"omitempty,min=8,max=50,password"`
	Age         int    `json:"age"         validate:"omitempty,gte=0,lte=200"`
}
