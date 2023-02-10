package entity

import (
	"github.com/blackmax1886/tas9-api/graph/model"
)

type User struct {
	ID    uint
	Name  string
	Email string
}

func NewUserFromEntity(user *User) *model.User {
	return &model.User{
		// ID:   fmt.Sprintf("%d", user.ID),
		Name:  user.Name,
		Email: user.Email,
	}
}
