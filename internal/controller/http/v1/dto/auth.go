package dto

type Login struct {
	Email    string `json:"email"    validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=50,password"`
}

type Token struct {
	Token string `json:"token"`
}
