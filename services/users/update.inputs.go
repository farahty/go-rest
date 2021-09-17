package usersService

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/google/uuid"
	"github.com/nimerfarahty/go-rest/models"
)

type UserUpdateInput struct {
	ID    uuid.UUID `json:"id"`
	Email *string   `json:"email"`
	Phone *string   `json:"phone"`
}

func (u *UserUpdateInput) ToModel() *models.User {

	return &models.User{
		Email: u.Email,
		Phone: u.Phone,
	}
}

func (u UserUpdateInput) Validate() error {
	return v.ValidateStruct(&u,
		v.Field(&u.Email, v.When(u.Email != nil, v.Required, is.Email)),
		v.Field(&u.Phone, v.When(u.Phone != nil, v.Required, v.Length(3, 20))),
	)
}
