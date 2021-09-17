package usersService

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/nimerfarahty/go-rest/models"
)

type CreateInput struct {
	Password *string `json:"password"`
	Email    *string `json:"email"`
	Phone    *string `json:"phone"`
}

func (u *CreateInput) ToModel() *models.User {

	return &models.User{
		Password: u.Password,
		Email:    u.Email,
		Phone:    u.Phone,
	}
}

func (u CreateInput) Validate() error {
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
