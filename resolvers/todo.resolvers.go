package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nimerfarahty/go-rest/models"
	todosService "github.com/nimerfarahty/go-rest/services/todo"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, todo models.CreateTodoInput) (*models.Todo, error) {
	return todosService.Create(todo)
}

func (r *queryResolver) Todos(ctx context.Context) ([]*models.Todo, error) {
	return todosService.Find()
}
