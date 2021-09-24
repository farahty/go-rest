package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nimerfarahty/go-rest/generated"
	"github.com/nimerfarahty/go-rest/models"
	usersService "github.com/nimerfarahty/go-rest/services/users"
)

func (r *mutationResolver) CreateUser(ctx context.Context, user models.CreateUserInput) (*models.User, error) {
	return usersService.CreateUser(&user)
}

func (r *mutationResolver) UpdateUser(ctx context.Context, user models.UpdateUserInput) (*models.User, error) {
	return usersService.Update(user)
}

func (r *queryResolver) Users(ctx context.Context) ([]*models.User, error) {
	return usersService.FindAll()
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
