package todosService

import (
	"github.com/nimerfarahty/go-rest/database"
	"github.com/nimerfarahty/go-rest/models"
)

func Create(input models.CreateTodoInput) (*models.Todo, error) {
	todo := &models.Todo{
		Name:      input.Name,
		Completed: input.Completed,
		UserID:    input.UserID,
	}

	if tx := database.Conn.Create(&todo); tx.Error != nil {
		return nil, tx.Error
	}

	return todo, nil
}

func Find() ([]*models.Todo, error) {

	todos := []*models.Todo{}

	if tx := database.Conn.Find(&todos); tx.Error != nil {
		return todos, tx.Error
	}

	return todos, nil
}
