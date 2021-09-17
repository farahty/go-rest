package authService

import (
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nimerfarahty/go-rest/models"
)

type LoginInput struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}
type LogoutInput struct {
	AccessToken string `json:"accessToken"`
}
type Tokens struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type LoginRespose struct {
	Tokens Tokens      `json:"token"`
	User   models.User `json:"user"`
}

func (l *LoginInput) Validate() error {
	return v.ValidateStruct(&l, v.Field(&l.Identity, v.Required), v.Field(&l.Password, v.Required))
}

func (l *LogoutInput) Validate() error {
	return v.ValidateStruct(&l, v.Field(&l.AccessToken, v.Required))
}

func (l *Tokens) Validate() error {
	return v.ValidateStruct(&l, v.Field(&l.AccessToken, v.Required), v.Field(&l.RefreshToken, v.Required))
}
