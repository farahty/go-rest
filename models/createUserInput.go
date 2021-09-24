package models

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (u *CreateUserInput) ToModel() *User {

	return &User{
		Password: u.Password,
		Email:    u.Email,
		Phone:    u.Phone,
	}
}

func (u CreateUserInput) Validate() error {
	return v.ValidateStruct(&u,
		v.Field(
			&u.Password,
			v.When(u.Password != nil,
				v.Required,
				v.Length(3, 50),
			),
		),
		v.Field(
			&u.Email,
			v.When(u.Phone == nil,
				v.Required,
				is.Email,
			),
		),
		v.Field(
			&u.Phone,
			v.When(
				u.Email == nil,
				v.Required,
				v.Length(3, 20),
			),
		),
	)
}
