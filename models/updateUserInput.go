package models

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (u *UpdateUserInput) ToModel() *User {

	return &User{
		Email:    u.Email,
		Phone:    u.Phone,
		Password: u.Password,
	}
}

func (u UpdateUserInput) Validate() error {
	return v.ValidateStruct(&u,
		v.Field(&u.ID, v.Required),
		v.Field(&u.Email, v.When(u.Email != nil, v.Required, is.Email)),
		v.Field(&u.Phone, v.When(u.Phone != nil, v.Required, v.Length(3, 20))),
	)
}
