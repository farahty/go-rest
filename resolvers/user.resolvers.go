package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/nimerfarahty/go-rest/models"
	usersService "github.com/nimerfarahty/go-rest/services/users"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user models.CreateUserInput) (*models.User, error) {
	return usersService.CreateUser(&user)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, user models.UpdateUserInput) (*models.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return usersService.FindAll()
}
