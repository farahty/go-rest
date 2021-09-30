package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nimerfarahty/go-rest/generated"
	"github.com/nimerfarahty/go-rest/models"
	authService "github.com/nimerfarahty/go-rest/services/auth"
)

func (r *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.LoginResponse, error) {
	return authService.Login(&input)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
