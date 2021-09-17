package usersService

import (
	"github.com/nimerfarahty/go-rest/database"
	"github.com/nimerfarahty/go-rest/models"
)

func CreateUser(user *CreateInput) (*models.User, error) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	model := user.ToModel()
	tx := database.Conn.Create(model)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return model, nil
}
