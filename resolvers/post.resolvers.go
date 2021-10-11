package resolvers

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/nimerfarahty/go-rest/generated"
	"github.com/nimerfarahty/go-rest/models"
	postsService "github.com/nimerfarahty/go-rest/services/posts"
)

func (r *mutationResolver) CreatePost(ctx context.Context, post models.CreatePostInput) (*models.Post, error) {
	return postsService.CreatePost(post)
}

func (r *queryResolver) Posts(ctx context.Context) ([]*models.Post, error) {
	return postsService.FindAll()
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
