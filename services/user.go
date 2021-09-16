package services

import (
	v "github.com/go-ozzo/ozzo-validation"
)

type UserInput struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func (u UserInput) Validate() error {
	return v.ValidateStruct(&u,
		v.Field(&u.UserName, v.Required, v.Length(3, 100)),
		v.Field(&u.Email, v.Required, v.Length(3, 100)),
	)
}

func CreateUser(user *UserInput) error {
	return user.Validate()
}
