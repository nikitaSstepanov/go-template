package converter

import (
	"app/internal/controller/http/v1/dto"
	"app/internal/entity"
)

func DtoUser(user *entity.User) *dto.Account {
	return &dto.Account{
		Id:       user.Id,
		Email:    user.Email,
		Name:     user.Name,
		Age:      user.Age,
		Verified: user.Verified,
	}
}

func EntityCreate(create dto.CreateUser) *entity.User {
	return &entity.User{
		Email:    create.Email,
		Name:     create.Name,
		Password: create.Password,
		Age:      create.Age,
	}
}

func EntityUpdate(update dto.UpdateUser) *entity.User {
	return &entity.User{
		Email:    update.Email,
		Name:     update.Name,
		Password: update.Password,
		Age:      update.Age,
	}
}
