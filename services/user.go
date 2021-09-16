package services

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

type UserCreateInput struct {
	Password string `json:"password"`
	UserName string `json:"userName"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

func (u UserCreateInput) Validate() error {
	return v.ValidateStruct(&u,
		v.Field(
			&u.Password,
			v.Required,
			v.Length(3, 50),
		),
		v.Field(
			&u.UserName,
			v.When(
				u.Email == "" && u.Phone == "",
				v.Required,
				v.Length(3, 50),
			),
		),
		v.Field(
			&u.Email,
			v.When(u.UserName == "" && u.Phone == "",
				v.Required,
				v.Length(3, 100),
				is.Email,
			),
		),
		v.Field(
			&u.Phone,
			v.When(
				u.Email == "" && u.UserName == "",
				v.Required, v.Length(3, 100),
			),
		),
	)
}

type UserUpdateInput struct {
	UserName string `json:"userName"`
	Email    string `json:"email"`
}

func (u UserUpdateInput) Validate() error {
	return v.ValidateStruct(&u,
		v.Field(&u.UserName, v.Required.When(u.Email == ""), v.Length(3, 100)),
		v.Field(&u.Email, v.Required, v.Length(3, 100)),
	)
}

func CreateUser(user *UserCreateInput) error {
	return user.Validate()
}
