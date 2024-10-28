package converter

import (
	"github.com/nikitaSstepanov/templates/golang/internal/controller/http/v1/dto"
	"github.com/nikitaSstepanov/templates/golang/internal/entity"
)

func EntityLogin(login dto.Login) *entity.User {
	return &entity.User{
		Email:    login.Email,
		Password: login.Password,
	}
}

func DtoToken(tokens *entity.Tokens) dto.Token {
	return dto.Token{
		Token: tokens.Access,
	}
}
