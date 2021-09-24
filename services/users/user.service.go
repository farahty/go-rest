package usersService

import (
	"github.com/nimerfarahty/go-rest/database"
	"github.com/nimerfarahty/go-rest/models"
)

func CreateUser(user *models.CreateUserInput) (*models.User, error) {

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

func Update(input models.UpdateUserInput) (*models.User, error) {

	if err := input.Validate(); err != nil {
		return nil, err
	}

	user := &models.User{}

	tx := database.Conn.First(&user, input.ID).Updates(input.ToModel())

	if tx.Error != nil {
		return nil, tx.Error
	}

	return user, nil

}

func FindAll() ([]*models.User, error) {

	users := []*models.User{}

	tx := database.Conn.Find(&users)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}
