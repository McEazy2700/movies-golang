package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	model "github.com/mceazy2700/movies-golang/db/models"
	"github.com/mceazy2700/movies-golang/graph/generated"
)

// AddMovie is the resolver for the addMovie field.
func (r *mutationResolver) AddMovie(ctx context.Context, title string, imdbCode string) (*model.MovieSuccess, error) {
	panic(fmt.Errorf("not implemented: AddMovie - addMovie"))
}

// UpdateViewsCount is the resolver for the updateViewsCount field.
func (r *mutationResolver) UpdateViewsCount(ctx context.Context, movieID string) (*model.MovieSuccess, error) {
	panic(fmt.Errorf("not implemented: UpdateViewsCount - updateViewsCount"))
}

// PopularMovies is the resolver for the popularMovies field.
func (r *queryResolver) PopularMovies(ctx context.Context) ([]*model.Movie, error) {
	panic(fmt.Errorf("not implemented: PopularMovies - popularMovies"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
